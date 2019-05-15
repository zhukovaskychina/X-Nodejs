#include "Util.h"

//using namespace std;

pid_t getProcessPidByName (string procname)
{
const char *proc_name=procname.c_str()
      FILE *fp;
      char buf[100];
      char cmd[200] = {'\0'};
      pid_t pid = -1;
      sprintf(cmd, "pidof %s", proc_name);

      if((fp = popen(cmd, "r")) != NULL)
      {
          if(fgets(buf, 255, fp) != NULL)
          {
              pid = atoi(buf);
          }
      }

      printf("pid = %d \n", pid);

      pclose(fp);
      return pid;
 }