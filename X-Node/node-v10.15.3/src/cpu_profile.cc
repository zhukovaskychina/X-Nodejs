
#include "v8.h"
#include "node.h"
#include "env.h"
#include "env-inl.h"
#include "cpu_profile.h"
#include "cpu_profile_node.h"
#include <iostream>


namespace node{
    using namespace std;
    namespace profileSpace{
        using node::Environment;
        using v8::Array;
        using v8::CpuProfile;
        using v8::CpuProfileNode;
        using v8::EscapableHandleScope;
        using v8::Handle;
        using v8::Number;
        using v8::HandleScope;
        using v8::Integer;
        using v8::Local;
        using v8::Object;
        using v8::ObjectTemplate;
        using v8::External;
        using v8::FunctionTemplate;
        using v8::String;
        using v8::Function;
        using v8::Value;
        using v8::Context;
        using v8::MaybeLocal;
        using v8::JSON;

        using namespace profileSpace;


        uint32_t uid_counter;

        Profile::Profile(){

            uid_counter=0;

        }

        Profile::~Profile(){
        //    uid_counter= null;

        }





        std::string Profile::New(const v8::CpuProfile *node) {

            Isolate *isolate = Isolate::GetCurrent();
            EscapableHandleScope scope(isolate);

            uid_counter++;

            Local<Object> profile=Object::New(isolate);

            const uint32_t uid_length = (((sizeof uid_counter) * 8) + 2)/3 + 2;

            char _uid[uid_length];

            sprintf(_uid, "%d", uid_counter);

            Local<Value> CPU = String::NewFromUtf8(isolate,"CPU");
            Local<Value> uid = String::NewFromUtf8(isolate,_uid);

            Local<String> title = node->GetTitle();
            if (!title->Length()) {
                char _title[8 + uid_length];
                sprintf(_title, "Profile %i", uid_counter);
                title = String::NewFromUtf8(isolate,_title);
            }
            Local<Value> head = ProfilerNode::New(node->GetTopDownRoot());

            profile->Set(String::NewFromUtf8(isolate,"typeId"),    CPU);
            profile->Set(String::NewFromUtf8(isolate,"uid"),       uid);
            profile->Set(String::NewFromUtf8(isolate,"title"),     title);
            profile->Set(String::NewFromUtf8(isolate,"head"),      head);



            Local<Value> start_time = Number::New(isolate,node->GetStartTime()/1000000);
            Local<Value> end_time = Number::New(isolate,node->GetEndTime()/1000000);


            uint32_t count = node->GetSamplesCount();
            Local<Array> timestamps = Array::New(isolate,count);
            Local<Array> samples = Array::New(isolate,count);
            for (uint32_t index = 0; index < count; ++index) {
                samples->Set(index, Integer::New(isolate,node->GetSample(index)->GetNodeId()));
                timestamps->Set(index, Number::New(isolate,static_cast<double>(node->GetSampleTimestamp(index))));
            }

            profile->Set(String::NewFromUtf8(isolate,"startTime"),   start_time);
            profile->Set(String::NewFromUtf8(isolate,"endTime"),     end_time);
            profile->Set(String::NewFromUtf8(isolate,"samples"),     samples);
            profile->Set(String::NewFromUtf8(isolate,"timestamps"),  timestamps);



            Local<Object> profiles=Object::New(isolate);
            profiles->Set(uid, profile);


            Local <Context> context = isolate->GetCurrentContext();
            MaybeLocal <String> jsonString = JSON::Stringify(context, scope.Escape(profile));
            std::string result;
            if (!jsonString.IsEmpty()) {
                Local <String> localString = jsonString.ToLocalChecked();
                result = *String::Utf8Value(localString);

            }else{
                result="";
            }
            return result;

        }
    }

}