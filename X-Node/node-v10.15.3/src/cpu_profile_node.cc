//
// Created by zhukovasky on 19-5-18.
//


#include "node.h"
#include "v8.h"

#include "cpu_profile_node.h"
#include "env.h"
#include <iostream>
namespace node {

    namespace profileSpace{
        using namespace v8;
        using namespace profileSpace;
        using v8::CpuProfileNode;
        using v8::Handle;
        using v8::String;
        using v8::Number;
        using v8::Integer;
        using v8::Value;
        using v8::Local;
        using v8::Object;
        using v8::Array;


        uint32_t ProfilerNode::UIDCounter = 1;

        v8::Local <v8::Value> ProfilerNode::GetLineTick_(const CpuProfileNode *node) {
            Isolate *isolate = Isolate::GetCurrent();
            v8::EscapableHandleScope scope(isolate);
            uint32_t count = node->GetHitLineCount();

            v8::CpuProfileNode::LineTick *entries = new v8::CpuProfileNode::LineTick[count];

            bool result = node->GetLineTicks(entries, count);

            Local <Value> lineTicks;
            if (result) {
                Local <Array> array = Array::New(isolate,count);
                for (uint32_t index = 0; index < count; index++) {
                    Local <Object> tick = Object::New(isolate);
                    tick->Set(String::NewFromUtf8(isolate,"line"), Integer::New(isolate,entries[index].line));
                    tick->Set(String::NewFromUtf8(isolate,"hitCount"), Integer::New(isolate,entries[index].hit_count));
                    array->Set(index, tick);
                }
                lineTicks = array;
            } else {
                lineTicks = v8::Null(isolate);
            }

            delete[] entries;

            return scope.Escape(lineTicks);
        }


        Local <Value> ProfilerNode::New(const CpuProfileNode *node) {
            Isolate *isolate = Isolate::GetCurrent();
            v8::EscapableHandleScope scope(isolate);
            int32_t count = node->GetChildrenCount();
            Local <Object> profile_node = Object::New(isolate);
            Local <Array> children = Array::New(isolate,count);

            for (int32_t index = 0; index < count; index++) {
                children->Set(index, ProfilerNode::New(node->GetChild(index)));
            }
            profile_node->Set(String::NewFromUtf8(isolate,"functionName"), node->GetFunctionName());
            profile_node->Set(String::NewFromUtf8(isolate,"url"), node->GetScriptResourceName());
            profile_node->Set(String::NewFromUtf8(isolate,"lineNumber"), Integer::New(isolate,node->GetLineNumber()));
            profile_node->Set(String::NewFromUtf8(isolate,"callUID"), Number::New(isolate,node->GetCallUid()));
            profile_node->Set(String::NewFromUtf8(isolate,"bailoutReason"), String::NewFromUtf8(isolate,node->GetBailoutReason()));
            profile_node->Set(String::NewFromUtf8(isolate,"id"), Integer::New(isolate,node->GetNodeId()));
            profile_node->Set(String::NewFromUtf8(isolate,"scriptId"), Integer::New(isolate,node->GetScriptId()));
            profile_node->Set(String::NewFromUtf8(isolate,"hitCount"), Integer::New(isolate,node->GetHitCount()));
            profile_node->Set(String::NewFromUtf8(isolate,"columnNumber"), Integer::New(isolate,node->GetColumnNumber()));
            profile_node->Set(String::NewFromUtf8(isolate,"children"), children);

            Local <Value> lineTicks = GetLineTick_(node);
            if (!lineTicks->IsNull()) {
                profile_node->Set(String::NewFromUtf8(isolate,"lineTicks"), lineTicks);
            }

            return scope.Escape(profile_node);
        }


    }

}