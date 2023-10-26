"use strict";(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[9895],{79895:function(t,e,a){a.r(e),a.d(e,{default:function(){return m}});var s=function(){var t=this,e=t._self._c;return e("el-drawer",{staticClass:"systemSet",attrs:{visible:t.dialogVisible,"show-close":!1,direction:"rtl","custom-class":"taskEditDrawer","close-on-press-escape":!1,"append-to-body":!0,"close-on-click-modal":!1,"before-close":t.handleClose,size:"90%"},on:{"update:visible":function(e){t.dialogVisible=e}}},[t.dialogVisible?e("div",[e("div",{attrs:{slot:"title"},slot:"title"}),e("el-row",{staticStyle:{"margin-bottom":"5px"}},[e("el-popover",{attrs:{placement:"bottom-start",width:"460",trigger:"click"},on:{show:function(e){return t.$refs["iconSelect"].reset()}}},[e("IconSelect",{ref:"iconSelect",on:{selected:t.iconSelected}}),e("template",{slot:"reference"},[e("span",{staticStyle:{position:"absolute",top:"0",left:"0"}},[t.serverData.icon?e("svg-icon",{staticClass:"iconfont",staticStyle:{"font-size":"23px",transform:"translateY(7px)"},attrs:{"icon-class":t.serverData.icon&&t.Base64.decode(t.serverData.icon)||""}}):e("i",{staticClass:"el-icon-paperclip",staticStyle:{"font-size":"23px",color:"white",transform:"translateY(7px)"}})],1)])],2),e("el-col",{staticStyle:{"padding-left":"26px"},attrs:{span:12}},[e("el-button",{attrs:{type:"text"}},[t._v(" "+t._s(t.serverData.name))]),e("el-button",{attrs:{type:"text"}},[t._v(t._s(!0===t.serverData.dialect?"私有":"公有"))]),e("el-button",{attrs:{type:"text"}},[t._v(t._s(t.serverData.status))]),e("el-button",{attrs:{type:"text"}},[t._v(t._s(t.dateTime(t.serverData.task_at)))]),e("el-button",{attrs:{type:"text"}},[t._v(t._s(t.dateTime(t.serverData.uptime)))])],1),e("el-col",{staticStyle:{"text-align":"right"},attrs:{span:12}},[e("i",{staticClass:"fa fa-code",staticStyle:{"margin-left":"10px",color:"#3d95ff",cursor:"pointer","font-weight":"bold"},attrs:{title:"显示/隐藏代码"},on:{click:t.editCode}}),e("i",{staticClass:"fa fa-refresh",staticStyle:{"margin-left":"10px",color:"#63a35c",cursor:"pointer"},attrs:{title:"刷新"},on:{click:t.refreshCode}}),e("el-popconfirm",{attrs:{title:"确定继续？"},on:{confirm:t.onSubmit}},["startup"!==t.serverData.name?e("i",{staticClass:"fa fa-cloud-upload",staticStyle:{"margin-left":"10px",color:"#3d95ff",cursor:"pointer"},attrs:{slot:"reference",title:"提交"},slot:"reference"}):t._e()]),e("i",{staticClass:"fa fa-undo",staticStyle:{"margin-left":"10px",color:"#3d95ff",cursor:"pointer"},attrs:{title:"重启"},on:{click:t.reloadService}}),e("i",{staticClass:"fa fa-close",staticStyle:{margin:"0 5px 0 10px",color:"#9e9e9e",cursor:"pointer"},attrs:{title:"关闭"},on:{click:function(e){t.dialogVisible=!1}}})],1)],1),e("el-row",[e("el-col",{attrs:{span:19}},[e("el-row",[t.codeShow?e("el-col",{attrs:{span:24}},[e("codeEditor",{staticStyle:{"min-height":"calc(100vh - 200px)"},model:{value:t.dialogValue,callback:function(e){t.dialogValue=e},expression:"dialogValue"}})],1):t._e()],1),e("el-row",{staticStyle:{"margin-top":"10px"}},[e("el-col",{attrs:{span:24}},[e("idlTools",{staticClass:"idlTools",attrs:{minionId:t.minionId}}),e("EventTask",{attrs:{dataTask:t.dataTask}})],1)],1)],1),e("el-col",{staticStyle:{"padding-left":"8px"},attrs:{span:5}},[e("el-card",[e("el-row",[e("el-col",{staticStyle:{"margin-bottom":"10px"}},[e("h3",{staticStyle:{color:"#fff"}},[t._v("任务信息")])]),e("el-col",{staticStyle:{"margin-bottom":"8px"}},[t._v(" 任务名称："+t._s(t.serverData.name)+" ")]),e("el-col",{staticStyle:{"margin-bottom":"8px"}},[t._v(" 外链详情："+t._s(t.nullChange(t.serverData.link))+" ")]),e("el-col",{staticStyle:{"margin-bottom":"8px"}},[t._v(" 本地哈希："+t._s(t.serverData.legal_hash)+" ")]),e("el-col",{staticStyle:{"margin-bottom":"8px"}},[t._v(" 上报哈希："+t._s(t.serverData.actual_hash)+" ")]),e("el-col",{staticStyle:{"margin-bottom":"8px"}},[t._v(" 报错信息："+t._s(t.nullChange(t.serverData.cause))+" ")]),e("el-col",{staticStyle:{"margin-bottom":"20px"}},[t._v(" 描述信息："+t._s(t.nullChange(t.serverData.desc))+" ")])],1),e("el-row",[e("el-col",{staticStyle:{"margin-bottom":"10px"}},[e("h3",{staticStyle:{color:"#fff"}},[t._v("服务信息")])]),t._l(t.serverData.runners,(function(a,s){return e("div",{key:s},[e("el-col",{staticStyle:{"margin-bottom":"5px"}},[e("span",{attrs:{type:"info",size:"mini"}},[t._v("名称："+t._s(a.name))])]),e("el-col",{staticStyle:{"margin-bottom":"5px"}},[e("span",{attrs:{type:"info",size:"mini"}},[t._v("类型："),""!==a.type?e("span",[t._v(t._s(a.type))]):t._e()])]),e("el-col",{staticStyle:{"margin-bottom":"15px"}},[e("span",{attrs:{type:"info",size:"mini"}},[t._v("状态："+t._s(a.status))])])],1)}))],2)],1)],1)],1)],1):t._e()])},i=[],o=a(91938),l=a(99815),r=a(78790);let c=a(19575).Base64;var n={name:"codeEdit",components:{IconSelect:o.Z,idlTools:l.Z,codeEditor:r.Z,EventTask:()=>a.e(8330).then(a.bind(a,68330))},props:{},data(){return{Base64:c,serverData:{},dialogVisible:!1,dialogValue:"",dataTask:{},codeShow:!0,minionId:"",initData:null}},methods:{reloadService(){var t={id:this.dataTask.minion_id,substance_id:this.serverData.id,dialect:this.serverData.dialect};this.$request.fetchPatchReload(t).then((t=>{this.$message.success("重启成功！！"),this.loadData(this.initData)})).catch((t=>{this.$message.error(t.data)}))},dateTime(t){var e=a(30381);return e(t).format("YYYY-MM-DD HH:mm:ss")},refreshCode(){this.loadData(this.initData)},editCode(){this.codeShow=!this.codeShow},nullChange(t){return""!==t&&" "!==t?t:"无"},handleClose(){},async iconSelected(t){try{this.serverData.icon=c.encode(t),await this.savaTask({id:this.serverData.id,icon:this.serverData.icon,chunk:this.serverData.chunk,version:this.serverData.version}),this.$emit("handleEdit",this.dataTask),this.$message.success("切换icon成功"),this.loadData(this.initData)}catch(e){this.$message.error(e||"操作出错")}},async savaTask(t){let e;return e=this.initData.dialect?()=>this.$request.fetchUpdateDialect(t):()=>this.$request.fetchUpdataCode(t),e()},async onSubmit(){try{let e=c.encode(this.dialogValue);var t={id:this.serverData.id,icon:this.serverData.icon,chunk:e,version:this.serverData.version};await this.savaTask(t),this.$message({message:"编辑成功",type:"success"}),this.$emit("handleEdit",this.dataTask),this.loadData(this.initData)}catch(e){this.$message.error(e)}},async loadData({dialect:t,minionId:e,taskId:a=null}){try{let t;t=e?()=>this.$request.fetchdetailDialect(e,a):()=>this.$request.fetchUpdataCodeEdit(a),this.minionId=e;const{data:s}=await t()||{};this.dataTask={minion_id:e,name:s.name,id:s.id},this.dialogValue=c.decode(s.chunk||" "),this.serverData=s||{}}catch(s){this.$message.error(s||"查询出错")}},async openTaskDetail(t){try{this.initData=t,await this.loadData(t),this.dialogVisible=!0}catch(e){this.$message.error(e||"查询出错")}}}},d=n,h=a(1001),p=(0,h.Z)(d,s,i,!1,null,"ec68c5be",null),m=p.exports}}]);