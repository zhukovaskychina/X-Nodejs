#include "node.h"
#include "env.h"
#include "cpu_profile.h"
#include "node_cpu_profile.h"


namespace node{

    namespace profileSpace{

        using v8::CpuProfile;
        using v8::Handle;
        using v8::Local;
        using v8::Object;
        using v8::Array;
        using v8::String;
        using v8::Value;
        using v8::Isolate;
        using v8::Boolean;
        using v8::Isolate;
        using v8::HandleScope;
        v8::CpuProfiler* current_cpu_profiler;

        const CpuProfile* profile;

        CpuProfiler::CpuProfiler() {
            Isolate *isolate=Isolate::GetCurrent();

            current_cpu_profiler= v8::CpuProfiler::New(isolate);
        }

        CpuProfiler::~CpuProfiler() {
            if(profile!= nullptr){
                //static_cast<CpuProfiler*>(profile)->Delete();
                delete profile;
            }
            if(current_cpu_profiler!= nullptr){
               // current_cpu_profiler->Dispose();
                 current_cpu_profiler= nullptr;
                 free(current_cpu_profiler);
            }
        }

        void CpuProfiler::StartProfiling() {

            Isolate *isolate=Isolate::GetCurrent();
            HandleScope scope(isolate);
            Local<String> title = String::NewFromUtf8(isolate,"");
            current_cpu_profiler->StartProfiling(title, true);

        }


        std::string CpuProfiler::StopProfiling() {


            Isolate *isolate=Isolate::GetCurrent();
            HandleScope scope(isolate);
            Local<String> title = String::NewFromUtf8(isolate,"");
            profile = current_cpu_profiler->StopProfiling(title);
            profileSpace::Profile *profileObject=new profileSpace::Profile();
            std::string result=profileObject->New(profile);
            delete profileObject;
            return result;


        }

        void CpuProfiler::SetSamplingInterval() {

            current_cpu_profiler->SetSamplingInterval(1000);
        }

        void CpuProfiler::Delete(){
            //profile->Delete();

        }
    }
    }
