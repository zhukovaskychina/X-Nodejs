#ifndef SRC_CPU_PROFILE_NODE_H_
#define SRC_CPU_PROFILE_NODE_H_

#include "v8-profiler.h"
#include "v8.h"
namespace node {
    using namespace v8;
    namespace profileSpace{
        class ProfilerNode{

        public:
            static v8::Local<v8::Value> New(const v8::CpuProfileNode* node);

            static uint32_t UIDCounter;

            static v8::Local<v8::Value> GetLineTick_(const v8::CpuProfileNode *node);
        };
    }

}


#endif