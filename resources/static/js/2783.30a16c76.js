(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[2783],{72783:function(t,e,i){"use strict";i.r(e),i.d(e,{default:function(){return Y}});var s=function(){var t=this,e=t._self._c;return e("el-row",{attrs:{gutter:10}},[e("el-col",{attrs:{span:15}},[e("el-row",[e("el-col",{attrs:{span:24}},[e("Riskdata")],1)],1),e("el-row",{staticStyle:{"margin-top":"10px"},attrs:{gutter:12}},[e("el-col",{attrs:{span:15}},[e("el-row",[e("el-col",{attrs:{span:24}},[e("Terminaldata")],1)],1),e("el-row",[e("el-col",{attrs:{span:24}},[e("el-row",{staticStyle:{"margin-top":"10px"}},[e("el-col",{attrs:{span:24}},[e("Clientversion")],1)],1)],1)],1)],1),e("el-col",{attrs:{span:9}},[e("el-card",{staticStyle:{height:"613px"},attrs:{shadow:"always"}},[e("Hosttop")],1)],1)],1)],1),e("el-col",{attrs:{span:9}},[e("Eventnotice")],1)],1)},a=[],n=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("el-row",[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard"},[t._v("风险数据")])]),e("el-col",{attrs:{span:12}},[e("span",{staticClass:"herfTitle"},[t._v("详情")])])],1),e("el-row",[e("el-col",{attrs:{span:12}},[e("div",{style:{float:"left",width:"100%",height:"200px"},attrs:{id:"risklvlEchart"}})]),e("el-col",{attrs:{span:12}},[e("div",{style:{float:"left",width:"100%",height:"200px"},attrs:{id:"riskstsEchart"}})])],1)],1)},o=[],l=i(83779),r={data(){return{risklvlData:{},riskstsData:{}}},created(){this.initRisk()},methods:{initRisk(){this.$request.getRisklvl().then((t=>{this.risklvlData=t.data,this.initRisklvl()})),this.$request.getRisksts().then((t=>{this.riskstsData=t.data,this.initRisksts()}))},initRisklvl(){let t=[],e=[];Math.max(...Object.values(this.risklvlData||{}));const i=Object.entries(this.risklvlData||{});this.risklvlData=this.risklvlData||[];for(let[n,o]of i){let i="",s="";"critical"===n?(i="紧急",s="#f86c6b"):"high"===n?(i="高风险",s="#fd8c00"):"middle"===n?(i="中风险",s="#ffc107"):"low"===n&&(i="低风险",s="#4dbd74"),t.push(i),e.push({value:Math.log(o),itemStyle:{color:s}})}const s={color:["#f86c6b","#fd8c00","#ffc107","#4dbd74"],tooltip:{show:!1},xAxis:{data:t},yAxis:{axisLine:{show:!1},axisTick:{show:!1},axisLabel:{show:!1},splitLine:{show:!1}},series:[{name:"风险等级",type:"bar",data:e,label:{show:!0,formatter:({dataIndex:t})=>i[t]?.[1],position:"top"}}]},a=l.init(document.getElementById("risklvlEchart"));a.setOption(s),window.addEventListener("resize",(()=>{a.resize()}))},initRisksts(){var t=[];for(let o in this.riskstsData){var e={};"ignore"===o?(e={name:"已忽略",value:this.riskstsData[o]},t.push(e)):"processed"===o?(e={name:"已处理",value:this.riskstsData[o]},t.push(e)):"unprocessed"===o&&(e={name:"未处理",value:this.riskstsData[o]},t.push(e))}let i=0;t.forEach((function(t){i+=t.value}));let s=["#f56c6c","#67c23a","#f7ba1e"];var a={backgroundColor:"#fff",tooltip:{},title:{text:i,subtext:"事件总数",x:"30%",y:"40%",textStyle:{fontSize:14,fontWeight:"normal",color:"#333"},subtextStyle:{fontSize:14,fontWeight:"normal",align:"center",color:"#555"}},legend:{orient:"vertical",right:"5%",align:"left",top:"middle",itemWidth:10,itemHeight:13,icon:"circle",itemGap:20,data:t,formatter:function(e){let s=0;for(var a="",n=0;n<t.length;n++)t[n].name==e&&(s=(t[n].value/i*100).toFixed(2)+"%",a+="{a|"+e+"}| {b|"+s+"}{b|"+t[n].value+"}");return a},textStyle:{color:"#000",rich:{a:{fontSize:12,padding:[0,10,0,6],color:"#666",fontFamily:"PingFang SC",fontStyle:"normal",lineHeight:20},b:{fontSize:12,padding:[0,10,0,6],color:"#666",fontFamily:"PingFang SC",fontStyle:"normal",lineHeight:20}}}},series:[{type:"pie",zlevel:3,radius:["50","90"],center:["35%","50%"],color:s,itemStyle:{normal:{borderWidth:10,borderColor:"#fff"}},data:t,labelLine:{normal:{show:!1}},label:{show:!1}}]};const n=l.init(document.getElementById("riskstsEchart"));n.setOption(a),window.addEventListener("resize",(()=>{n.resize()}))}}},c=r,h=i(81656),u=(0,h.A)(c,n,o,!1,null,null,null),d=u.exports,f=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("el-row",[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard"},[t._v("终端数据")])]),e("el-col",{attrs:{span:12}},[e("span",{staticClass:"herfTitle",on:{click:t.detailClick}},[t._v("详情")])])],1),e("div",{staticClass:"statusChart",style:{float:"left",width:"100%",height:"249px"},attrs:{id:"statusChart"}})],1)},p=[],v={name:"Terminaldata",data(){return{statusData:{}}},created(){this.initStatus()},methods:{initStatus(){this.$request.getStatus().then((t=>{this.statusData=t.data,this.initChart()})).catch((()=>{this.statusData=[]}))},initChart(){var t=[];for(let o in this.statusData){var e={};"online"===o?(e={name:"在线",value:this.statusData[o]},t.push(e)):"offline"===o?(e={name:"离线",value:this.statusData[o]},t.push(e)):"inactive"===o?(e={name:"未激活",value:this.statusData[o]},t.push(e)):"deleted"===o&&(e={name:"已删除",value:this.statusData[o]},t.push(e))}let i=0;t.forEach((function(t){i+=t.value}));let s=["#67c23a","#909399","#e6a23c","#f56c6c"];var a={backgroundColor:"#fff",tooltip:{},title:{text:i,subtext:"终端总数",x:"30%",y:"40%",textStyle:{fontSize:14,fontWeight:"normal",color:"#333"},subtextStyle:{fontSize:14,fontWeight:"normal",align:"center",color:"#555"}},legend:{orient:"vertical",right:"5%",align:"left",top:"middle",itemWidth:10,itemHeight:13,icon:"circle",itemGap:20,data:t,formatter:function(e){let s=0;for(var a="",n=0;n<t.length;n++)t[n].name==e&&(s=(t[n].value/i*100).toFixed(2)+"%",a+="{a|"+e+"}| {b|"+s+"}{b|"+t[n].value+"}");return a},textStyle:{color:"#000",rich:{a:{fontSize:12,padding:[0,10,0,6],color:"#666",fontFamily:"PingFang SC",fontStyle:"normal",lineHeight:20},b:{fontSize:12,padding:[0,10,0,6],color:"#666",fontFamily:"PingFang SC",fontStyle:"normal",lineHeight:20}}}},series:[{type:"pie",zlevel:3,radius:["50","90"],center:["35%","50%"],color:s,itemStyle:{normal:{borderWidth:10,borderColor:"#fff"}},data:t,labelLine:{normal:{show:!1}},label:{show:!1}}]};const n=l.init(document.getElementById("statusChart"));n.setOption(a),window.addEventListener("resize",(()=>{n.resize()}))},detailClick(){this.$router.push("/node")}}},m=v,g=(0,h.A)(m,f,p,!1,null,null,null),w=g.exports,y=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("el-row",[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard"},[t._v("客户端版本")])]),e("el-col",{attrs:{span:12}},[e("span",{staticClass:"herfTitle"},[t._v("详情")])])],1),e("div",{staticClass:"client",style:{float:"left",width:"100%",height:"253px"},attrs:{id:"client"}})],1)},S=[],x={name:"Terminaldata",data(){return{editionData:[]}},created(){this.initEdition()},methods:{initEdition(){this.$request.getEdtition().then((t=>{this.editionData=t.data.data?.slice(0,10)||[],this.mychartes()})).catch((()=>{this.editionData=[]}))},mychartes(){var t=this.editionData.map((t=>t.edition)),e=this.editionData.map((t=>0===t.total?0:1===t.total?.3:Math.log(t.total))),i={color:["#f86c6b","#fd8c00","#ffc107","#4dbd74"],tooltip:{show:!1},xAxis:{type:"category",inverse:!0,data:t,axisLabel:{interval:0}},yAxis:{show:!1},series:[{name:"数量",type:"bar",data:e,label:{show:!0,formatter:({dataIndex:t})=>this.editionData[t].total,position:"top"}}]};const s=l.init(document.getElementById("client"));s.setOption(i),window.addEventListener("resize",(()=>{s.resize()}))}}},_=x,b=(0,h.A)(_,y,S,!1,null,null,null),C=b.exports,P=function(){var t=this,e=t._self._c;return e("div",[e("el-row",[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard"},[t._v("风险主机TOP10")])]),e("el-col",{attrs:{span:12}},[e("span",{staticClass:"herfTitle"},[t._v("详情")])])],1),e("el-table",{staticClass:"continer",staticStyle:{width:"100%"},attrs:{data:t.tableData.records?t.tableData.records.slice(0,10):[]}},[e("el-table-column",{attrs:{type:"index",width:"50"},scopedSlots:t._u([{key:"default",fn:function({row:i,$index:s}){return[e("div",{class:s+1===1?"topCircular first":s+1===2?"topCircular second":s+1===3?"topCircular third":"topCircular"},[t._v(" "+t._s(s+1)+" ")])]}}])}),e("el-table-column",{attrs:{prop:"name",label:"终端",width:"200"}}),e("el-table-column",{attrs:{prop:"count",label:"漏洞数"}})],1)],1)},L=[],T={data(){return{tableData:[],loading:!1,currentPage:1,pageSize:15}},created(){this.init()},methods:{init(){var t={current:this.currentPage,size:this.pageSize,group:"inet",filters:this.filters};this.loading=!0,this.$request.fetchGroup(t).then((t=>{this.tableData=t.data,this.loading=!1}))}}},k=T,z=(0,h.A)(k,P,L,!1,null,null,null),D=z.exports,M=function(){var t=this,e=t._self._c;return e("div",{staticClass:"continer"},[e("el-card",{staticClass:"event"},[e("el-row",{staticStyle:{"margin-bottom":"10px"}},[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard"},[t._v("风险事件")])]),e("el-col",{attrs:{span:12}},[e("span",{staticClass:"herfTitle",on:{click:t.refresh}},[t._v("刷新")])])],1),e("vue-seamless-scroll",{staticClass:"warp",attrs:{data:t.riskData.records,"class-option":t.classOption}},t._l(t.riskData.records,(function(i,s){return e("div",{key:s,staticClass:"event2"},[e("el-row",[e("el-col",{attrs:{span:24}},["紧急"===i.level?e("div",{staticClass:"severity-critical-bg text-pull"},[e("svg",{staticClass:"icon",attrs:{fill:"#fff",t:"1658471401691",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2218"}},[e("path",{attrs:{d:"M906.496819 314.75159 767.929904 314.75159c-22.190395-38.46301-52.763714-71.749114-89.50143-96.65127l80.132038-80.132038-69.77618-69.77618L581.529563 175.444824c-22.190395-5.424546-45.613874-8.629541-69.529563-8.629541-23.916712 0-47.339168 3.204995-69.529563 8.629541L335.215668 68.191078l-69.77618 69.77618 80.132038 80.132038c-36.737716 24.902156-67.311035 58.18826-89.50143 96.65127L117.503181 314.750567l0 98.624205 103.062284 0c-2.711761 16.025998-4.438079 32.546254-4.438079 49.312102l0 49.312102-98.624205 0 0 98.624205 98.624205 0 0 49.312102c0 16.765849 1.726318 33.286104 4.438079 49.312102L117.503181 709.247386l0 98.624205 138.566916 0c51.03842 88.268346 146.456606 147.936307 255.929904 147.936307S716.891484 896.14096 767.929904 807.872614l138.566916 0 0-98.624205L803.434535 709.24841c2.711761-16.025998 4.438079-32.546254 4.438079-49.312102l0-49.312102 98.624205 0 0-98.624205-98.624205 0 0-49.312102c0-16.765849-1.726318-33.286104-4.438079-49.312102l103.062284 0L906.496819 314.75159zM610.624205 709.24841 413.375795 709.24841l0-98.624205 197.24841 0L610.624205 709.24841zM610.624205 512 413.375795 512l0-98.624205 197.24841 0L610.624205 512z","p-id":"2219"}})])]):t._e(),"高危"===i.level?e("div",{staticClass:"severity-high-bg text-pull"},[e("svg",{staticClass:"icon",attrs:{fill:"#fff",t:"1658471401691",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2218"}},[e("path",{attrs:{d:"M906.496819 314.75159 767.929904 314.75159c-22.190395-38.46301-52.763714-71.749114-89.50143-96.65127l80.132038-80.132038-69.77618-69.77618L581.529563 175.444824c-22.190395-5.424546-45.613874-8.629541-69.529563-8.629541-23.916712 0-47.339168 3.204995-69.529563 8.629541L335.215668 68.191078l-69.77618 69.77618 80.132038 80.132038c-36.737716 24.902156-67.311035 58.18826-89.50143 96.65127L117.503181 314.750567l0 98.624205 103.062284 0c-2.711761 16.025998-4.438079 32.546254-4.438079 49.312102l0 49.312102-98.624205 0 0 98.624205 98.624205 0 0 49.312102c0 16.765849 1.726318 33.286104 4.438079 49.312102L117.503181 709.247386l0 98.624205 138.566916 0c51.03842 88.268346 146.456606 147.936307 255.929904 147.936307S716.891484 896.14096 767.929904 807.872614l138.566916 0 0-98.624205L803.434535 709.24841c2.711761-16.025998 4.438079-32.546254 4.438079-49.312102l0-49.312102 98.624205 0 0-98.624205-98.624205 0 0-49.312102c0-16.765849-1.726318-33.286104-4.438079-49.312102l103.062284 0L906.496819 314.75159zM610.624205 709.24841 413.375795 709.24841l0-98.624205 197.24841 0L610.624205 709.24841zM610.624205 512 413.375795 512l0-98.624205 197.24841 0L610.624205 512z","p-id":"2219"}})])]):t._e(),"中危"===i.level?e("div",{staticClass:"severity-medium-bg text-pull"},[e("svg",{staticClass:"icon",attrs:{fill:"#fff",t:"1658471401691",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2218"}},[e("path",{attrs:{d:"M906.496819 314.75159 767.929904 314.75159c-22.190395-38.46301-52.763714-71.749114-89.50143-96.65127l80.132038-80.132038-69.77618-69.77618L581.529563 175.444824c-22.190395-5.424546-45.613874-8.629541-69.529563-8.629541-23.916712 0-47.339168 3.204995-69.529563 8.629541L335.215668 68.191078l-69.77618 69.77618 80.132038 80.132038c-36.737716 24.902156-67.311035 58.18826-89.50143 96.65127L117.503181 314.750567l0 98.624205 103.062284 0c-2.711761 16.025998-4.438079 32.546254-4.438079 49.312102l0 49.312102-98.624205 0 0 98.624205 98.624205 0 0 49.312102c0 16.765849 1.726318 33.286104 4.438079 49.312102L117.503181 709.247386l0 98.624205 138.566916 0c51.03842 88.268346 146.456606 147.936307 255.929904 147.936307S716.891484 896.14096 767.929904 807.872614l138.566916 0 0-98.624205L803.434535 709.24841c2.711761-16.025998 4.438079-32.546254 4.438079-49.312102l0-49.312102 98.624205 0 0-98.624205-98.624205 0 0-49.312102c0-16.765849-1.726318-33.286104-4.438079-49.312102l103.062284 0L906.496819 314.75159zM610.624205 709.24841 413.375795 709.24841l0-98.624205 197.24841 0L610.624205 709.24841zM610.624205 512 413.375795 512l0-98.624205 197.24841 0L610.624205 512z","p-id":"2219"}})])]):t._e(),"低危"===i.level?e("div",{staticClass:"severity-low-bg text-pull"},[e("svg",{staticClass:"icon",attrs:{fill:"#fff",t:"1658471401691",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2218"}},[e("path",{attrs:{d:"M906.496819 314.75159 767.929904 314.75159c-22.190395-38.46301-52.763714-71.749114-89.50143-96.65127l80.132038-80.132038-69.77618-69.77618L581.529563 175.444824c-22.190395-5.424546-45.613874-8.629541-69.529563-8.629541-23.916712 0-47.339168 3.204995-69.529563 8.629541L335.215668 68.191078l-69.77618 69.77618 80.132038 80.132038c-36.737716 24.902156-67.311035 58.18826-89.50143 96.65127L117.503181 314.750567l0 98.624205 103.062284 0c-2.711761 16.025998-4.438079 32.546254-4.438079 49.312102l0 49.312102-98.624205 0 0 98.624205 98.624205 0 0 49.312102c0 16.765849 1.726318 33.286104 4.438079 49.312102L117.503181 709.247386l0 98.624205 138.566916 0c51.03842 88.268346 146.456606 147.936307 255.929904 147.936307S716.891484 896.14096 767.929904 807.872614l138.566916 0 0-98.624205L803.434535 709.24841c2.711761-16.025998 4.438079-32.546254 4.438079-49.312102l0-49.312102 98.624205 0 0-98.624205-98.624205 0 0-49.312102c0-16.765849-1.726318-33.286104-4.438079-49.312102l103.062284 0L906.496819 314.75159zM610.624205 709.24841 413.375795 709.24841l0-98.624205 197.24841 0L610.624205 709.24841zM610.624205 512 413.375795 512l0-98.624205 197.24841 0L610.624205 512z","p-id":"2219"}})])]):t._e(),"无风险"===i.level?e("div",{staticClass:"severity-no-bg text-pull"},[e("svg",{staticClass:"icon",attrs:{fill:"#fff",t:"1658471401691",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"2218"}},[e("path",{attrs:{d:"M906.496819 314.75159 767.929904 314.75159c-22.190395-38.46301-52.763714-71.749114-89.50143-96.65127l80.132038-80.132038-69.77618-69.77618L581.529563 175.444824c-22.190395-5.424546-45.613874-8.629541-69.529563-8.629541-23.916712 0-47.339168 3.204995-69.529563 8.629541L335.215668 68.191078l-69.77618 69.77618 80.132038 80.132038c-36.737716 24.902156-67.311035 58.18826-89.50143 96.65127L117.503181 314.750567l0 98.624205 103.062284 0c-2.711761 16.025998-4.438079 32.546254-4.438079 49.312102l0 49.312102-98.624205 0 0 98.624205 98.624205 0 0 49.312102c0 16.765849 1.726318 33.286104 4.438079 49.312102L117.503181 709.247386l0 98.624205 138.566916 0c51.03842 88.268346 146.456606 147.936307 255.929904 147.936307S716.891484 896.14096 767.929904 807.872614l138.566916 0 0-98.624205L803.434535 709.24841c2.711761-16.025998 4.438079-32.546254 4.438079-49.312102l0-49.312102 98.624205 0 0-98.624205-98.624205 0 0-49.312102c0-16.765849-1.726318-33.286104-4.438079-49.312102l103.062284 0L906.496819 314.75159zM610.624205 709.24841 413.375795 709.24841l0-98.624205 197.24841 0L610.624205 709.24841zM610.624205 512 413.375795 512l0-98.624205 197.24841 0L610.624205 512z","p-id":"2219"}})])]):t._e(),e("div",{staticStyle:{"font-size":"12px","margin-top":"-5px"}},[e("b",[t._v(" "+t._s(i.level))]),e("span",{staticClass:"rowTitle"},[e("b",[t._v("远程IP:")])]),t._v(" "+t._s(i.remote_ip)+" "),e("span",{staticClass:"rowTitle"},[e("b",[t._v("主题:")])]),t._v(" "+t._s(i.subject)+" "),e("span",{staticClass:"rowTitle"},[e("b",[t._v("时间：")]),e("span",[t._v(t._s(i.created_at&&t.moment(i.created_at).format("YYYY-MM-DD HH:mm:ss")))])])])])],1)],1)})),0)],1)],1)},E=[],H=i(7866),W=i.n(H),A=i(95093),F=i.n(A),O={name:"Eventnotice",components:{vueSeamlessScroll:W()},data(){return{moment:F(),currentPage:1,pageSize:100,riskData:{},filters:[],classOption:{step:.5},timer:null}},created(){this.getEventData()},beforeDestroy(){clearTimeout(this.timer)},methods:{refresh(){this.getEventData()},getEventData(){clearTimeout(this.timer);var t={current:this.currentPage,size:this.pageSize,filters:this.filters};this.riskLoading=!0,this.$request.fetchRisks(t).then((t=>{this.riskData=t.data,this.riskLoading=!1,setTimeout((()=>{this.getEventData()}),6e4)}))}}},$=O,R=(0,h.A)($,M,E,!1,null,"ab998338",null),B=R.exports,j={name:"index",components:{Riskdata:d,Terminaldata:w,Clientversion:C,Hosttop:D,Eventnotice:B}},q=j,I=(0,h.A)(q,s,a,!1,null,"339bb1e6",null),Y=I.exports},7866:function(t){!function(e,i){t.exports=i()}("undefined"!=typeof self&&self,(function(){return function(t){function e(s){if(i[s])return i[s].exports;var a=i[s]={i:s,l:!1,exports:{}};return t[s].call(a.exports,a,a.exports,e),a.l=!0,a.exports}var i={};return e.m=t,e.c=i,e.d=function(t,i,s){e.o(t,i)||Object.defineProperty(t,i,{configurable:!1,enumerable:!0,get:s})},e.n=function(t){var i=t&&t.__esModule?function(){return t.default}:function(){return t};return e.d(i,"a",i),i},e.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},e.p="",e(e.s=1)}([function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),i(4)();var s=i(5),a=i(6);e.default={name:"vue-seamless-scroll",data:function(){return{xPos:0,yPos:0,delay:0,copyHtml:"",height:0,width:0,realBoxWidth:0}},props:{data:{type:Array,default:function(){return[]}},classOption:{type:Object,default:function(){return{}}}},computed:{leftSwitchState:function(){return this.xPos<0},rightSwitchState:function(){return Math.abs(this.xPos)<this.realBoxWidth-this.width},leftSwitchClass:function(){return this.leftSwitchState?"":this.options.switchDisabledClass},rightSwitchClass:function(){return this.rightSwitchState?"":this.options.switchDisabledClass},leftSwitch:function(){return{position:"absolute",margin:this.height/2+"px 0 0 -"+this.options.switchOffset+"px",transform:"translate(-100%,-50%)"}},rightSwitch:function(){return{position:"absolute",margin:this.height/2+"px 0 0 "+(this.width+this.options.switchOffset)+"px",transform:"translateY(-50%)"}},float:function(){return this.isHorizontal?{float:"left",overflow:"hidden"}:{overflow:"hidden"}},pos:function(){return{transform:"translate("+this.xPos+"px,"+this.yPos+"px)",transition:"all "+this.ease+" "+this.delay+"ms",overflow:"hidden"}},defaultOption:function(){return{step:1,limitMoveNum:5,hoverStop:!0,direction:1,openTouch:!0,singleHeight:0,singleWidth:0,waitTime:1e3,switchOffset:30,autoPlay:!0,navigation:!1,switchSingleStep:134,switchDelay:400,switchDisabledClass:"disabled",isSingleRemUnit:!1}},options:function(){return a({},this.defaultOption,this.classOption)},navigation:function(){return this.options.navigation},autoPlay:function(){return!this.navigation&&this.options.autoPlay},scrollSwitch:function(){return this.data.length>=this.options.limitMoveNum},hoverStopSwitch:function(){return this.options.hoverStop&&this.autoPlay&&this.scrollSwitch},canTouchScroll:function(){return this.options.openTouch},isHorizontal:function(){return this.options.direction>1},baseFontSize:function(){return this.options.isSingleRemUnit?parseInt(window.getComputedStyle(document.documentElement,null).fontSize):1},realSingleStopWidth:function(){return this.options.singleWidth*this.baseFontSize},realSingleStopHeight:function(){return this.options.singleHeight*this.baseFontSize},step:function(){var t=this.options.step;return this.isHorizontal?this.realSingleStopWidth:this.realSingleStopHeight,t}},methods:{reset:function(){this._cancle(),this._initMove()},leftSwitchClick:function(){if(this.leftSwitchState)return Math.abs(this.xPos)<this.options.switchSingleStep?void(this.xPos=0):void(this.xPos+=this.options.switchSingleStep)},rightSwitchClick:function(){if(this.rightSwitchState)return this.realBoxWidth-this.width+this.xPos<this.options.switchSingleStep?void(this.xPos=this.width-this.realBoxWidth):void(this.xPos-=this.options.switchSingleStep)},_cancle:function(){cancelAnimationFrame(this.reqFrame||"")},touchStart:function(t){var e=this;if(this.canTouchScroll){var i=void 0,s=t.targetTouches[0],a=this.options,n=a.waitTime,o=a.singleHeight,l=a.singleWidth;this.startPos={x:s.pageX,y:s.pageY},this.startPosY=this.yPos,this.startPosX=this.xPos,o&&l?(i&&clearTimeout(i),i=setTimeout((function(){e._cancle()}),n+20)):this._cancle()}},touchMove:function(t){if(!(!this.canTouchScroll||t.targetTouches.length>1||t.scale&&1!==t.scale)){var e=t.targetTouches[0],i=this.options.direction;this.endPos={x:e.pageX-this.startPos.x,y:e.pageY-this.startPos.y},event.preventDefault();var s=Math.abs(this.endPos.x)<Math.abs(this.endPos.y)?1:0;1===s&&i<2?this.yPos=this.startPosY+this.endPos.y:0===s&&i>1&&(this.xPos=this.startPosX+this.endPos.x)}},touchEnd:function(){var t=this;if(this.canTouchScroll){var e=void 0,i=this.options.direction;if(this.delay=50,1===i)this.yPos>0&&(this.yPos=0);else if(0===i){var s=this.realBoxHeight/2*-1;this.yPos<s&&(this.yPos=s)}else if(2===i)this.xPos>0&&(this.xPos=0);else if(3===i){var a=-1*this.realBoxWidth;this.xPos<a&&(this.xPos=a)}e&&clearTimeout(e),e=setTimeout((function(){t.delay=0,t._move()}),this.delay)}},enter:function(){this.hoverStopSwitch&&this._stopMove()},leave:function(){this.hoverStopSwitch&&this._startMove()},_move:function(){this.isHover||(this._cancle(),this.reqFrame=requestAnimationFrame(function(){var t=this,e=this.realBoxHeight/2,i=this.realBoxWidth/2,s=this.options,a=s.direction,n=s.waitTime,o=this.step;1===a?(Math.abs(this.yPos)>=e&&(this.$emit("ScrollEnd"),this.yPos=0),this.yPos-=o):0===a?(this.yPos>=0&&(this.$emit("ScrollEnd"),this.yPos=-1*e),this.yPos+=o):2===a?(Math.abs(this.xPos)>=i&&(this.$emit("ScrollEnd"),this.xPos=0),this.xPos-=o):3===a&&(this.xPos>=0&&(this.$emit("ScrollEnd"),this.xPos=-1*i),this.xPos+=o),this.singleWaitTime&&clearTimeout(this.singleWaitTime),this.realSingleStopHeight?Math.abs(this.yPos)%this.realSingleStopHeight<o?this.singleWaitTime=setTimeout((function(){t._move()}),n):this._move():this.realSingleStopWidth&&Math.abs(this.xPos)%this.realSingleStopWidth<o?this.singleWaitTime=setTimeout((function(){t._move()}),n):this._move()}.bind(this)))},_initMove:function(){var t=this;this.$nextTick((function(){var e=t.options.switchDelay,i=t.autoPlay,s=t.isHorizontal;if(t._dataWarm(t.data),t.copyHtml="",s){t.height=t.$refs.wrap.offsetHeight,t.width=t.$refs.wrap.offsetWidth;var a=t.$refs.slotList.offsetWidth;i&&(a=2*a+1),t.$refs.realBox.style.width=a+"px",t.realBoxWidth=a}if(!i)return t.ease="linear",void(t.delay=e);t.ease="ease-in",t.delay=0,t.scrollSwitch?(t.copyHtml=t.$refs.slotList.innerHTML,setTimeout((function(){t.realBoxHeight=t.$refs.realBox.offsetHeight,t._move()}),0)):(t._cancle(),t.yPos=t.xPos=0)}))},_dataWarm:function(t){t.length},_startMove:function(){this.isHover=!1,this._move()},_stopMove:function(){this.isHover=!0,this.singleWaitTime&&clearTimeout(this.singleWaitTime),this._cancle()}},mounted:function(){this._initMove()},watch:{data:function(t,e){this._dataWarm(t),s(t,e)||this.reset()},autoPlay:function(t){t?this.reset():this._stopMove()}},beforeCreate:function(){this.reqFrame=null,this.singleWaitTime=null,this.isHover=!1,this.ease="ease-in"},beforeDestroy:function(){this._cancle(),clearTimeout(this.singleWaitTime)}}},function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=i(2),a=function(t){return t&&t.__esModule?t:{default:t}}(s);a.default.install=function(t){var e=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};t.component(e.componentName||a.default.name,a.default)},"undefined"!=typeof window&&window.Vue&&Vue.component(a.default.name,a.default),e.default=a.default},function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=i(0),a=i.n(s);for(var n in s)"default"!==n&&function(t){i.d(e,t,(function(){return s[t]}))}(n);var o=i(7),l=i(3),r=l(a.a,o.a,!1,null,null,null);e.default=r.exports},function(t,e){t.exports=function(t,e,i,s,a,n){var o,l=t=t||{},r=typeof t.default;"object"!==r&&"function"!==r||(o=t,l=t.default);var c,h="function"==typeof l?l.options:l;if(e&&(h.render=e.render,h.staticRenderFns=e.staticRenderFns,h._compiled=!0),i&&(h.functional=!0),a&&(h._scopeId=a),n?(c=function(t){t=t||this.$vnode&&this.$vnode.ssrContext||this.parent&&this.parent.$vnode&&this.parent.$vnode.ssrContext,t||"undefined"==typeof __VUE_SSR_CONTEXT__||(t=__VUE_SSR_CONTEXT__),s&&s.call(this,t),t&&t._registeredComponents&&t._registeredComponents.add(n)},h._ssrRegister=c):s&&(c=s),c){var u=h.functional,d=u?h.render:h.beforeCreate;u?(h._injectStyles=c,h.render=function(t,e){return c.call(e),d(t,e)}):h.beforeCreate=d?[].concat(d,c):[c]}return{esModule:o,exports:l,options:h}}},function(t,e){var i=function(){window.cancelAnimationFrame=function(){return window.cancelAnimationFrame||window.webkitCancelAnimationFrame||window.mozCancelAnimationFrame||window.oCancelAnimationFrame||window.msCancelAnimationFrame||function(t){return window.clearTimeout(t)}}(),window.requestAnimationFrame=function(){return window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||window.oRequestAnimationFrame||window.msRequestAnimationFrame||function(t){return window.setTimeout(t,1e3/60)}}()};t.exports=i},function(t,e){var i=function(t,e){if(t===e)return!0;if(t.length!==e.length)return!1;for(var i=0;i<t.length;++i)if(t[i]!==e[i])return!1;return!0};t.exports=i},function(t,e){function i(){Array.isArray||(Array.isArray=function(t){return"[object Array]"===Object.prototype.toString.call(t)});var t=void 0,e=void 0,a=void 0,n=void 0,o=void 0,l=void 0,r=1,c=arguments[0]||{},h=!1,u=arguments.length;if("boolean"==typeof c&&(h=c,c=arguments[1]||{},r++),"object"!==(void 0===c?"undefined":s(c))&&"function"!=typeof c&&(c={}),r===u)return c;for(;r<u;r++)if(null!=(e=arguments[r]))for(t in e)a=c[t],n=e[t],o=Array.isArray(n),h&&n&&("object"===(void 0===n?"undefined":s(n))||o)?(o?(o=!1,l=a&&Array.isArray(a)?a:[]):l=a&&"object"===(void 0===a?"undefined":s(a))?a:{},c[t]=i(h,l,n)):void 0!==n&&(c[t]=n);return c}var s="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t};t.exports=i},function(t,e,i){"use strict";var s=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{ref:"wrap"},[t.navigation?i("div",{class:t.leftSwitchClass,style:t.leftSwitch,on:{click:t.leftSwitchClick}},[t._t("left-switch")],2):t._e(),t._v(" "),t.navigation?i("div",{class:t.rightSwitchClass,style:t.rightSwitch,on:{click:t.rightSwitchClick}},[t._t("right-switch")],2):t._e(),t._v(" "),i("div",{ref:"realBox",style:t.pos,on:{mouseenter:t.enter,mouseleave:t.leave,touchstart:t.touchStart,touchmove:t.touchMove,touchend:t.touchEnd}},[i("div",{ref:"slotList",style:t.float},[t._t("default")],2),t._v(" "),i("div",{style:t.float,domProps:{innerHTML:t._s(t.copyHtml)}})])])},a=[],n={render:s,staticRenderFns:a};e.a=n}]).default}))}}]);