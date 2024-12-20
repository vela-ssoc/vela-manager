"use strict";(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[5753],{86389:function(t,e,a){a.d(e,{A:function(){return d}});var i=function(){var t=this,e=t._self._c;return e("detail",t._g(t._b({},"detail",t.$attrs,!1),t.$listeners))},n=[],l=a(24869),s={components:{detail:l.A}},o=s,r=a(81656),c=(0,r.A)(o,i,n,!1,null,null,null),d=c.exports},95854:function(t,e,a){a.r(e),a.d(e,{default:function(){return E}});var i=function(){var t=this,e=t._self._c;return e("div",[t.chartShow?e("Recent",{on:{getDay:t.getDay}}):t._e(),e("el-card",[t.eventShow?e("Event",{attrs:{dayTime:t.dayTime},on:{activeVulnClick:t.activeVulnClick,activechartShow:t.activechartShow}}):t._e(),t.attackShow?e("Attack",{attrs:{dayTime:t.dayTime},on:{activeVulnClick:t.activeVulnClick,activechartShow:t.activechartShow}}):t._e()],1)],1)},n=[],l=a(92161),s=(a(12878),function(){var t=this,e=t._self._c;return e("span",[e("el-row",[e("el-col",{attrs:{span:18}},[e("filterOption",{ref:"selectSearchRef",attrs:{options:t.condData.conditions||[]},on:{change:t.searchBasic},scopedSlots:t._u([{key:"textSlot",fn:function(){return[t._v(" 添加过滤条件 ")]},proxy:!0}]),model:{value:t.filters,callback:function(e){t.filters=e},expression:"filters"}})],1),e("el-col",{staticStyle:{"text-align":"right"},attrs:{span:6}},[e("el-button",{attrs:{type:"success",plain:"",size:"mini"},on:{click:t.echartShow}},[t._v(t._s(!0===t.topEchartShow?"隐藏图表":"显示图表"))]),e("el-button-group",[e("el-button",{attrs:{type:"info",size:"mini"},on:{click:t.eventShow}},[t._v("登录事件")]),e("el-button",{attrs:{type:"primary",size:"mini"}},[t._v("攻击统计")])],1)],1)],1),e("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.attackLoading,expression:"attackLoading"}],staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:t.attackData.records,"header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[e("el-table-column",{attrs:{prop:"addr",label:"尝试地址"},scopedSlots:t._u([{key:"default",fn:function(a){return[" 0.0.0.0"===a.row.addr||"127.0.0.1"===a.row.addr||""===a.row.addr||"nil"===a.row.addr?e("span",[t._v("本地")]):e("span",[t._v(t._s(a.row.addr))])]}}])}),e("el-table-column",{attrs:{prop:"msg",label:"描述"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.msg)+" ")]}}])}),e("el-table-column",{attrs:{prop:"count",label:"总计"}}),e("el-table-column",{attrs:{label:"操作",width:"150"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.evnetDetail(a.row)}}},[t._v("登录历史")])]}}])})],1),e("Page",{attrs:{size:t.attackData.size,current:t.attackData.current,pageTotal:t.attackData.total},on:{handleSizeChange:t.handleSizeChange,handleCurrentChange:t.handleCurrentChange}}),e("el-drawer",{attrs:{title:"登录历史",visible:t.detailShow,direction:"rtl",size:"90%"},on:{"update:visible":function(e){t.detailShow=e}}},[e("div",[e("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.detalLoading,expression:"detalLoading"}],attrs:{data:t.detailData.records,"header-cell-style":{color:"#606266",background:"#f5f7fa"},stripe:"",align:"center"}},[e("el-table-column",{attrs:{prop:"inet",label:"终端IP"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.innetactive(a.row)}}},[t._v(t._s(a.row.inet))])]}}])}),e("el-table-column",{attrs:{prop:"user",label:"账户"}}),e("el-table-column",{attrs:{prop:"msg",label:"描述"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.msg)+" ")]}}])}),e("el-table-column",{attrs:{prop:"addr",label:"登录地址"},scopedSlots:t._u([{key:"default",fn:function(a){return[" 0.0.0.0"===a.row.addr||"127.0.0.1"===a.row.addr||""===a.row.addr||"nil"===a.row.addr?e("span",[t._v("本地")]):e("span",[t._v(t._s(a.row.addr))])]}}])}),e("el-table-column",{attrs:{prop:"logon_at",label:"登录时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t._f("date")(e.row.logon_at,"yyyy-MM-dd hh:mm:ss"))+" ")]}}])}),e("el-table-column",{attrs:{label:"操作",width:"90"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.loginAlert(a.row)}}},[t._v("告警")]),e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.loginIgnore(a.row)}}},[t._v("忽略")])]}}])})],1),e("Page",{attrs:{size:t.detailData.size,current:t.detailData.current,pageTotal:t.detailData.total},on:{handleSizeChange:t.detailhandleSizeChange,handleCurrentChange:t.detailhandleCurrentChange}})],1)]),e("el-drawer",{attrs:{title:t.ipTtitel.inet,visible:t.drawer,direction:t.directions,"before-close":t.drawerClose,size:"90%"},on:{"update:visible":function(e){t.drawer=e}}},[e("node-detail",{ref:"ipDetail",attrs:{minionId:t.ipTtitel.minion_id}})],1)],1)}),o=[],r=a(86389),c=a(72953),d=a(95093),h=a.n(d),u={components:{Page:l.A,NodeDetail:r.A,filterOption:c.A},props:["dayTime"],name:"vuln",data(){return{condData:{},value:"",operator:"",field:"",filters:[{key:"logon_at",operator:"gt",value:h()().startOf("days").subtract(15,"days"),clearable:!1}],attackLoading:!1,topEchartShow:!1,attackData:{},selectForm:{name:"",inet:"",msg:[]},msgArray:[{value:"登录成功",label:"登录成功"},{value:"用户注销",label:"用户注销"},{value:"登录失败",label:"登录失败"}],fileDetailData:{},fileDetailLoading:!1,startTime:null,endTime:null,ipTtitel:{},drawer:!1,directions:"rtl",currentPage:1,pageSize:15,pageTotal:0,searchData:"",vulnDetailShow:!1,vulnDetailData:{},vulnLoading:!1,activeIndex:null,activeDetailData:{},detailShow:!1,detalLoading:!1,detailcurrentPage:1,detailpageSize:15,detailData:{}}},created(){this.initGetAccontLoginList(),this.getCond()},methods:{echartShow(){this.topEchartShow=!this.topEchartShow,this.$emit("activechartShow",this.topEchartShow)},drawerClose(){this.drawer=!1,this.ipTtitel={}},innetactive(t){this.ipTtitel=t,this.drawer=!0},loginAlert(t){this.$request.fetchLoginAlert({id:t.id}).then((t=>{this.$message({message:"已告警！！！",type:"success"}),this.evnetDetail()}))},loginIgnore(t){this.$request.fetchLoginIgnore({id:t.id}).then((t=>{this.$message({message:"已忽略！！！",type:"success"}),this.evnetDetail()}))},detailhandleSizeChange(t){this.detailpageSize=t,this.evnetDetail(this.activeDetailData)},detailhandleCurrentChange(t){this.detailcurrentPage=t,this.evnetDetail(this.activeDetailData)},vulnmDetailCancel(){this.detailShow=!1,this.activeDetailData={}},evnetDetail(t){this.activeDetailData=t;const e=[{key:"addr",value:this.activeDetailData.addr,operator:"eq"},{key:"msg",value:this.activeDetailData.msg,operator:"eq"},this.filters[0]],a={current:this.detailcurrentPage,size:this.detailpageSize,filters:e};this.$request.fetchAccontLoginList(a).then((t=>{this.detailData=t.data,this.detailShow=!0}))},getCond(){this.$request.fetchCond().then((t=>{this.condData=t.data}))},eventShow(){this.$emit("activeVulnClick",!0,!1)},initGetAccontLoginList(){this.attackLoading=!0;var t={current:this.currentPage,size:this.pageSize,filters:this.filters};this.$request.fetchLogonAttack(t).then((t=>{this.attackData=t.data})).catch((t=>{this.$message.error(t)})).finally((()=>{this.attackLoading=!1}))},searchBasic(){this.currentPage=1,this.initGetAccontLoginList()},handleSizeChange(t){this.pageSize=t,this.initGetAccontLoginList()},handleCurrentChange(t){this.currentPage=t,this.initGetAccontLoginList()}},watch:{dayTime:function(t,e){this.dayTime=t,this.initGetAccontLoginList()}}},p=u,g=a(81656),f=(0,g.A)(p,s,o,!1,null,"60c96427",null),v=f.exports,m=function(){var t=this,e=t._self._c;return e("span",[e("el-row",[e("el-col",{attrs:{span:18}},[e("filterOption",{ref:"selectSearchRef",attrs:{options:t.condData.conditions||[]},on:{change:t.searchBasic},scopedSlots:t._u([{key:"textSlot",fn:function(){return[t._v(" 添加过滤条件 ")]},proxy:!0}]),model:{value:t.filters,callback:function(e){t.filters=e},expression:"filters"}})],1),e("el-col",{staticStyle:{"text-align":"right"},attrs:{span:6}},[e("el-button",{attrs:{type:"success",plain:"",size:"mini"},on:{click:t.echartShow}},[t._v(t._s(!0===t.topEchartShow?"隐藏图表":"显示图表"))]),e("el-button",{attrs:{type:"primary",plain:"",size:"mini"},on:{click:t.sheetode}},[t._v("导出"),e("i",{staticClass:"el-icon-upload2 el-icon--right"})]),e("el-button-group",[e("el-button",{attrs:{type:"primary",size:"mini"}},[t._v("登录事件")]),e("el-button",{attrs:{type:"info",size:"mini"},on:{click:t.accountShow}},[t._v("攻击统计")])],1)],1)],1),e("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.accountLoading,expression:"accountLoading"}],staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:t.accountData.records,"header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[e("el-table-column",{attrs:{prop:"inet",label:"终端IP"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.innetactive(a.row)}}},[t._v(" "+t._s(a.row.inet))]),e("i",{staticClass:"el-icon-circle-plus-outline",staticStyle:{"margin-left":"5px",color:"#3d95ff",cursor:"pointer"},on:{click:function(e){return t.searchAdd({key:"inet",desc:"节点 IPv4"},a.row.inet)}}})]}}])}),e("el-table-column",{attrs:{prop:"user",label:"账户"}}),e("el-table-column",{attrs:{prop:"msg",label:"描述"},scopedSlots:t._u([{key:"default",fn:function(a){return[t._v(" "+t._s(a.row.msg)),e("i",{staticClass:"el-icon-circle-plus-outline",staticStyle:{"margin-left":"5px",color:"#3d95ff",cursor:"pointer"},on:{click:function(e){return t.searchAdd({key:"msg",desc:"登录描述"},a.row.msg)}}})]}}])}),e("el-table-column",{attrs:{prop:"addr",label:"登录地址"},scopedSlots:t._u([{key:"default",fn:function(a){return[" 0.0.0.0"===a.row.addr||"127.0.0.1"===a.row.addr||""===a.row.addr||"nil"===a.row.addr?e("span",[t._v("本地")]):e("span",[t._v(t._s(a.row.addr))])]}}])}),e("el-table-column",{attrs:{prop:"logon_at",label:"登录时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t._f("date")(e.row.logon_at,"yyyy-MM-dd hh:mm:ss"))+" ")]}}])}),e("el-table-column",{attrs:{label:"操作",width:"90"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.loginAlert(a.row)}}},[t._v("告警")]),e("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){return t.loginIgnore(a.row)}}},[t._v("忽略")])]}}])})],1),e("Page",{attrs:{size:t.accountData.size,current:t.accountData.current,pageTotal:t.accountData.total},on:{handleSizeChange:t.handleSizeChange,handleCurrentChange:t.handleCurrentChange}}),e("el-drawer",{staticClass:"vulnDrawer",attrs:{visible:t.fileDetailShow,direction:"rtl",size:"40%"},on:{"update:visible":function(e){t.fileDetailShow=e},close:t.fileDetailCancel}},[e("div",{directives:[{name:"loading",rawName:"v-loading",value:t.fileDetailLoading,expression:"fileDetailLoading"}],staticStyle:{padding:"0 20px"}},[e("span",{staticStyle:{"font-size":"13px",color:"#86909c","font-weight":"bold"}},[e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard",staticStyle:{"font-size":"12px"}},[t._v("文件信息")])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("文件名"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.vuln_id))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("大小"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.name))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("组件"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.version))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("修改"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.cvss_score))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("创建"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.purl))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("更新"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.title))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("哈希"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.title))])])],1),e("el-row",{staticStyle:{"margin-bottom":"14px"}},[e("el-col",{attrs:{span:24}},[t._v("节点"),e("span",{staticClass:"subtitle"},[t._v(t._s(t.vulnDetailData.title))])])],1)],1),e("el-row",{staticStyle:{"margin-top":"20px"}},[e("el-col",{attrs:{span:12}},[e("span",{staticClass:"controlCard",staticStyle:{"font-size":"12px"}},[t._v("相关文件")])])],1),e("el-table",{staticStyle:{"margin-left":"-10px"},attrs:{data:t.vulnDetailData.materials}},[e("el-table-column",{attrs:{prop:"filename",label:"文件名称"}}),e("el-table-column",{attrs:{prop:"modify_time",label:"修改时间",width:"150"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t._f("date")(e.row.modify_time,"yyyy-MM-dd hh:mm:ss"))+" ")]}}])})],1)],1)]),e("el-drawer",{staticClass:"inetDrawer",attrs:{title:t.ipTtitel.inet,visible:t.drawer,direction:t.directions,"before-close":t.drawerClose,size:"90%"},on:{"update:visible":function(e){t.drawer=e}}},[e("node-detail",{ref:"ipDetail",attrs:{minionId:t.ipTtitel.minion_id}})],1)],1)},y=[],w={components:{Page:l.A,NodeDetail:r.A,filterOption:c.A},props:["dayTime"],name:"vuln",data(){return{topEchartShow:!1,condData:{},value:"",operator:"",field:"",filters:[{key:"logon_at",operator:"gt",value:h()().startOf("days").subtract(15,"days"),clearable:!1}],accountLoading:!1,accountData:{},selectForm:{name:"",inet:"",msg:[]},msgArray:[{value:"登录成功",label:"登录成功"},{value:"用户注销",label:"用户注销"},{value:"登录失败",label:"登录失败"}],fileDetailShow:!1,fileDetailData:{},fileDetailLoading:!1,startTime:null,endTime:null,currentPage:1,pageSize:15,pageTotal:0,searchData:"",vulnDetailShow:!1,vulnDetailData:{},vulnLoading:!1,activeIndex:null,eventLogon:"logon_at",ipTtitel:{},drawer:!1,directions:"rtl"}},created(){this.initGetAccontLoginList(),this.getCond()},methods:{echartShow(){this.topEchartShow=!this.topEchartShow,this.$emit("activechartShow",this.topEchartShow)},drawerClose(){this.drawer=!1,this.ipTtitel={}},innetactive(t){this.ipTtitel=t,this.drawer=!0},sheetode(){var t="/api/v1/logon/csv?authorization="+localStorage.getItem("token");for(var e in this.filters)t=t+"&filters="+JSON.stringify(this.filters[e]);window.open(t)},searchAdd(t,e){this.currentPage=1;const a=this.$refs.selectSearchRef;a.pushFilter({key:t.key,operator:"eq",value:e}),this.initGetAccontLoginList()},getCond(){this.$request.fetchCond().then((t=>{this.condData=t.data}))},accountShow(){this.$emit("activeVulnClick",!1,!0)},initGetAccontLoginList(){this.accountLoading=!0;var t={current:this.currentPage,size:this.pageSize,filters:this.filters};this.$request.fetchAccontLoginList(t).then((t=>{this.accountLoading=!1,this.accountData=t.data})).catch((t=>{this.accountLoading=!1,this.$message.error(t)}))},fileDetailCancel(){this.fileDetailShow=!1,this.fileDetailData={}},loginIgnore(t){this.$request.fetchLoginIgnore({id:t.id}).then((t=>{this.$message({message:"已忽略！！！",type:"success"}),this.initGetAccontLoginList()}))},loginAlert(t){this.$request.fetchLoginAlert({id:t.id}).then((t=>{this.$message({message:"已告警！！！",type:"success"}),this.initGetAccontLoginList()}))},searchBasic(){this.currentPage=1,this.initGetAccontLoginList()},handleSizeChange(t){this.pageSize=t,this.initGetAccontLoginList()},handleCurrentChange(t){this.currentPage=t,this.initGetAccontLoginList()}},watch:{dayTime:function(t,e){this.dayTime=t,this.initGetAccontLoginList()}}},b=w,_=(0,g.A)(b,m,y,!1,null,"00ad7a0f",null),S=_.exports,D=function(){var t=this,e=t._self._c;return e("el-card",{staticStyle:{"margin-bottom":"10px"},attrs:{shadow:"never"}},[e("el-row",[e("el-col",{attrs:{span:24}},[e("div",{style:{float:"left",width:"100%",height:"240px"},attrs:{id:"riskstsEchart"}})]),e("div",[e("el-checkbox-group",{staticClass:"day_select",attrs:{size:"mini"},on:{change:t.dayActive},model:{value:t.dayCheck,callback:function(e){t.dayCheck=e},expression:"dayCheck"}},t._l(t.dayArray,(function(a,i){return e("div",{key:i},[e("el-checkbox",{attrs:{label:a.value}},[t._v(t._s(a.label))])],1)})),0)],1)],1)],1)},k=[],C=a(83779),L={data(){return{rencentData:{},date:[],failed:[],logout:[],success:[],selectDate:"15",dayCheck:["15"],dayArray:[{value:"1",label:"查询一天"},{value:"3",label:"查询三天"},{value:"5",label:"查询五天"},{value:"7",label:"查询一周"},{value:"15",label:"查询半月"},{value:"30",label:"查询每月"}]}},created(){this.initRecent(this.selectDate)},methods:{initRecent(t){this.$request.fetchLoginRecent(t).then((t=>{this.date=t.data.map((t=>t.date)),this.failed=t.data.map((t=>t.failed)),this.logout=t.data.map((t=>t.logout)),this.success=t.data.map((t=>t.success)),this.chartRecent()}))},dayActive(){this.dayCheck.length>1&&this.dayCheck.shift(),void 0===this.dayCheck[0]&&(this.dayCheck[0]="15"),this.$emit("getDay",this.dayCheck[0]),this.initRecent(this.dayCheck[0])},chartRecent(){var t={color:["#5470C6","#91CC75","#FAC858"],grid:{left:"4%",right:"8%",top:"16%",bottom:"8%"},tooltip:{trigger:"axis",axisPointer:{type:"cross",crossStyle:{color:"#999"}}},legend:{orient:"vertical",top:0,right:0,itemWidth:10,itemHeight:10,data:["登录失败","用户注销","登录成功"]},xAxis:{type:"category",data:this.date},yAxis:[{type:"value",show:!1}],series:[{backgroundStyle:{color:"rgba(111, 162, 135, 0.2)"},emphasis:{focus:"series"},name:"登录失败",type:"bar",data:this.failed},{backgroundStyle:{color:"rgba(111, 162, 135, 0.2)"},emphasis:{focus:"series"},name:"用户注销",type:"bar",data:this.logout},{backgroundStyle:{color:"rgba(111, 162, 135, 0.2)"},emphasis:{focus:"series"},name:"登录成功",type:"bar",data:this.success}]};const e=C.init(document.getElementById("riskstsEchart"));e.setOption(t),window.addEventListener("resize",(()=>{e.resize()}))}}},x=L,A=(0,g.A)(x,D,k,!1,null,"7b687987",null),z=A.exports,T={components:{Page:l.A,Attack:v,Event:S,Recent:z},name:"index",data(){return{eventShow:!0,chartShow:!1,attackShow:!1,dayTime:"15"}},methods:{activechartShow(){this.chartShow=!this.chartShow},activeVulnClick(t,e){this.chartShow=!1,this.eventShow=t,this.attackShow=e},getDay(t){this.dayTime=t}}},$=T,P=(0,g.A)($,i,n,!1,null,"1c1931ee",null),E=P.exports}}]);