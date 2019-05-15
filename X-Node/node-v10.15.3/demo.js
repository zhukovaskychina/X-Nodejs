var http = require("http");
const trace_events = require('trace_events');
const categories = ['node.perf', 'node.async_hooks','v8'];
const tracing = trace_events.createTracing({ categories });
tracing.enable();
// do stuff
http.createServer(function(req,res){
   res.writeHead(200,{"Content-type":"text/blain"});
 let a=1;   
for(var i=0;i<100000000000000;i++){
       a=i+1;
	//console.log(a);
   }   
res.write("Hello NodeJs"+a);
   res.end();

}).listen(8888);

console.log(trace_events.getEnabledCategories());
tracing.disable();

