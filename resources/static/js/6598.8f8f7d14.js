"use strict";(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[6598],{92161:function(e,t,i){i.d(t,{A:function(){return c}});var a=function(){var e=this,t=e._self._c;return t("div",{class:"事件详情"===e.nameShow?"pageClass":"",staticStyle:{height:"20px","padding-top":"10px"}},[t("span",{staticStyle:{display:"inline-block","font-size":"13px","min-width":"35.5px",height:"28px","line-height":"30px","vertical-align":"top","box-sizing":"border-box",float:"right"}},[e._v("共"+e._s(e.pageTotal)+"条")]),t("el-pagination",{staticStyle:{float:"right"},attrs:{background:"",layout:"jumper,sizes,prev, pager, next","page-size":e.size||15,"current-page":e.pageCurrent,total:e.pageTotal,"page-sizes":e.pageSizeArray||[15,30,45]},on:{"update:currentPage":function(t){e.pageCurrent=t},"update:current-page":function(t){e.pageCurrent=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)},s=[],r={name:"page",props:{pageSizeArray:Array,size:Number,pageTotal:Number,current:Number,nameShow:String},data(){return{pageCurrent:this.current}},methods:{handleSizeChange(e){this.$emit("handleSizeChange",e)},handleCurrentChange(e){this.$emit("handleCurrentChange",e)}},watch:{current:function(e){this.pageCurrent=e}}},n=r,o=i(81656),l=(0,o.A)(n,a,s,!1,null,"277403d7",null),c=l.exports},56598:function(e,t,i){i.r(t),i.d(t,{default:function(){return w}});var a=function(){var e=this,t=e._self._c;return t("div",{staticStyle:{width:"100%"}},[t("el-card",[t("el-row",[t("el-col",{attrs:{span:12}},[t("el-button",{staticStyle:{"margin-right":"7px"},attrs:{"el-button":"",type:"primary",size:"mini"},on:{click:e.addContact}},[e._v("添加 ")]),t("el-input",{staticClass:"input-with-select",staticStyle:{width:"20%","margin-right":"7px"},attrs:{placeholder:"请输入名称",size:"mini"},model:{value:e.selectInput,callback:function(t){e.selectInput=t},expression:"selectInput"}}),t("el-button",{staticStyle:{color:"#909399",background:"#f4f4f5","border-color":"#d3d4d6"},attrs:{icon:"el-icon-search",size:"mini"},on:{click:e.getContact}})],1)],1),t("el-table",{staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:e.contactData,"header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[t("el-table-column",{attrs:{prop:"name",label:"联系人",width:"120"}}),t("el-table-column",{attrs:{label:"风险类型"},scopedSlots:e._u([{key:"default",fn:function(i){return[i.row.risks?t("div",{staticStyle:{width:"95%",margin:"0 auto","text-align":"left"}},[e.ipRe(i.row.risks).length<5?t("span",e._l(e.ipRe(i.row.risks),(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i))])})),1):t("span",[e._l(e.ipRe(i.row.risks).slice(0,4),(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i))])})),t("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"Right Center 提示文字",placement:"right"}},[t("el-tag",{attrs:{size:"mini",type:"info"}},[e._v(" + "+e._s(e.ipRe(i.row.risks).length))]),t("div",{attrs:{slot:"content"},slot:"content"},e._l(e.ipRe(i.row.risks),(function(i,a){return t("span",{key:a,staticStyle:{"padding-right":"15px"}},[e._v(e._s(i)),t("br")])})),0)],1)],2)]):t("div",[e._v("-")])]}}])}),t("el-table-column",{attrs:{label:"事件类型"},scopedSlots:e._u([{key:"default",fn:function(i){return[i.row.events?t("div",{staticStyle:{width:"95%",margin:"0 auto","text-align":"left"}},[e.ipRe(i.row.events).length<5?t("span",e._l(e.ipRe(i.row.events),(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i))])})),1):t("span",[e._l(e.ipRe(i.row.events).slice(0,4),(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i))])})),t("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"Right Center 提示文字",placement:"right"}},[t("el-tag",{attrs:{size:"mini",type:"info"}},[e._v(" + "+e._s(e.ipRe(i.row.events).length))]),t("div",{attrs:{slot:"content"},slot:"content"},e._l(e.ipRe(i.row.events),(function(i,a){return t("span",{key:a,staticStyle:{"padding-right":"15px"}},[e._v(e._s(i)),t("br")])})),0)],1)],2)]):e._e()]}}])}),t("el-table-column",{attrs:{label:"方式"},scopedSlots:e._u([{key:"default",fn:function(i){return e._l(i.row.ways,(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i)+" ")])}))}}])}),t("el-table-column",{attrs:{label:"账号"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("div",{staticStyle:{width:"50%","text-align":"left"}},[""!==i.row.email?t("p",[t("el-tag",{attrs:{type:"info",size:"mini"}},[e._v("邮箱："+e._s(i.row.email))])],1):e._e(),""!==i.row.mobile?t("p",[t("el-tag",{attrs:{type:"info",size:"mini"}},[e._v("手机："+e._s(i.row.mobile))])],1):e._e(),""!==i.row.dong?t("p",[t("el-tag",{attrs:{type:"info",size:"mini"}},[e._v("咚咚："+e._s(i.row.dong))])],1):e._e()])]}}])}),t("el-table-column",{attrs:{label:"操作",width:"190"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-button",{attrs:{size:"mini"},on:{click:function(t){return e.handleEdit(i.$index,i.row)}}},[e._v("编辑 ")]),t("el-button",{attrs:{size:"mini"},on:{click:function(t){return e.handleIgnore(i.$index,i.row)}}},[e._v("忽略 ")]),t("el-popconfirm",{staticStyle:{"margin-left":"7px"},attrs:{title:"确定删除吗？"},on:{confirm:function(t){return e.subDelete(i.row.id)}}},[t("el-button",{attrs:{slot:"reference",type:"danger",icon:"el-icon-delete",size:"mini"},slot:"reference"})],1)]}}])})],1),t("Page",{attrs:{size:e.size,current:e.current,pageTotal:e.pageTotal,pageSizeArray:e.pageSizeArray},on:{handleSizeChange:e.handleSizeChange,handleCurrentChange:e.handleCurrentChange}})],1),t("ContactForm",{ref:"forms",attrs:{"is-add":e.isAdd}}),t("FormIgnore",{ref:"formsIgnore",attrs:{isAddIgnore:e.isAddIgnore}})],1)},s=[],r=i(92161),n=function(){var e=this,t=e._self._c;return t("el-dialog",{attrs:{"before-close":e.handleClose,"el-dialog":"","close-on-click-modal":!1,visible:e.dialogVisible,title:e.isAdd?"新增告警":"编辑告警","append-to-body":"",width:"650px"},on:{"update:visible":function(t){e.dialogVisible=t}}},[t("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,size:"small","label-width":"90px"}},[t("el-form-item",{attrs:{label:"联系人",prop:"name"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1),t("el-form-item",{attrs:{label:"事件类型"}},[t("el-input",{attrs:{type:"textarea",rows:10},model:{value:e.eventsData,callback:function(t){e.eventsData=t},expression:"eventsData"}})],1),t("el-form-item",{attrs:{label:"风险类型"}},[t("el-input",{attrs:{type:"textarea",rows:10},model:{value:e.risksData,callback:function(t){e.risksData=t},expression:"risksData"}})],1),t("el-form-item",{attrs:{label:"通知方式",prop:"ways"}},[t("el-select",{staticStyle:{width:"100%"},attrs:{multiple:""},model:{value:e.form.ways,callback:function(t){e.$set(e.form,"ways",t)},expression:"form.ways"}},e._l(e.waysList,(function(e,i){return t("el-option",{key:i,attrs:{value:e.value,label:e.label}})})),1)],1),t("el-form-item",{attrs:{label:"手机号",prop:"mobile"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.mobile,callback:function(t){e.$set(e.form,"mobile",t)},expression:"form.mobile"}})],1),t("el-form-item",{attrs:{label:"咚咚",prop:"dong"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.dong,callback:function(t){e.$set(e.form,"dong",t)},expression:"form.dong"}})],1),t("el-form-item",{attrs:{label:"邮箱",prop:"email"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.email,callback:function(t){e.$set(e.form,"email",t)},expression:"form.email"}})],1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:e.handleClose}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmit("form")}}},[e._v(e._s(e.isAdd?"创建":"修改"))])],1)],1)},o=[],l={name:"formRelation",props:{isAdd:Boolean},data(){return{form:{id:"",name:"",events:[],risks:[],ways:[],mobile:"",dong:"",email:""},waysList:[{label:"短信",value:"sms"},{label:"微信",value:"wechat"},{label:"咚咚",value:"dong"},{label:"邮箱",value:"email"}],eventsData:"",risksData:"",dialogVisible:!1,rules:{name:[{required:!0,message:"请输入联系人",trigger:"blur"}],ways:[{required:!0,message:"请选择通知方式",trigger:"change"}]}}},methods:{handleClose(){this.dialogVisible=!1,this.eventsData="",this.risksData="",this.form={}},mapArray(e){if(null!==e)return e.map((e=>e.id))},onSubmit(e){this.$refs[e].validate((e=>{this.form.events=this.eventsData.split("\n").filter((e=>""!==e)),this.form.risks=this.risksData.split("\n").filter((e=>""!==e)),e&&(this.isAdd?this.$request.fetchPostContact(this.form).then((()=>{this.$message({message:"添加成功!!!",type:"success"}),this.dialogVisible=!1,this.form={},this.$parent.getContact()})).catch((e=>{this.$message.error(e.data)})):this.$request.fetchPutContact(this.form).then((()=>{this.$message({message:"修改成功!!!",type:"success"}),this.dialogVisible=!1,this.form={},this.$parent.getContact()})).catch((e=>{this.$message.error(e.data)})))}))}}},c=l,d=i(81656),m=(0,d.A)(c,n,o,!1,null,"50045e50",null),u=m.exports,f=function(){var e=this,t=e._self._c;return t("el-dialog",{attrs:{"before-close":e.handleClose,"el-dialog":"","close-on-click-modal":!1,visible:e.ignoreShow,title:"忽略","append-to-body":"",width:"650px"},on:{"update:visible":function(t){e.ignoreShow=t}}},[t("el-form",{ref:"form",attrs:{model:e.form,size:"small","label-width":"90px"}},[t("el-form-item",{attrs:{label:"忽略事件"}},[t("el-input",{attrs:{type:"textarea",rows:10},model:{value:e.form.event_code,callback:function(t){e.$set(e.form,"event_code",t)},expression:"form.event_code"}})],1),t("el-form-item",{attrs:{label:"忽略风险"}},[t("el-input",{attrs:{type:"textarea",rows:10},model:{value:e.form.risk_code,callback:function(t){e.$set(e.form,"risk_code",t)},expression:"form.risk_code"}})],1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:e.handleClose}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmit("form")}}},[e._v("修改")])],1)],1)},h=[],p={name:"formIgnore",props:{isAddIgnore:Boolean},data(){return{form:{id:"",name:"",events:[],risks:[],ways:[],mobile:"",dong:"",email:"",event_code:"",risk_code:""},ignoreShow:!1}},methods:{handleClose(){this.ignoreShow=!1,this.form={}},mapArray(e){if(null!==e)return e.map((e=>e.id))},onSubmit(e){this.$refs[e].validate((e=>{if(e){let e=i(8127).Base64,a=e.encode(this.form.event_code),s=e.encode(this.form.risk_code);var t={id:this.form.id,name:this.form.name,events:this.form.events,risks:this.form.risks,ways:this.form.ways,mobile:this.form.mobile,dong:this.form.dong,email:this.form.email,event_code:a,risk_code:s};this.$request.fetchPutContact(t).then((()=>{this.$message({message:"修改成功!!!",type:"success"}),this.ignoreShow=!1,this.form={},this.$parent.getContact()})).catch((e=>{this.$message.error(e.data)}))}}))}}},g=p,b=(0,d.A)(g,f,h,!1,null,"1ccd4652",null),v=b.exports,y={name:"index",components:{ContactForm:u,Page:r.A,FormIgnore:v},data(){return{pageSizeArray:[10,20,30,40],current:1,size:10,pageTotal:0,contactData:[],isAdd:!0,isAddIgnore:!0,delLoading:!1,selectInput:""}},created(){this.getContact()},methods:{ipRe(e){if(null!==e){var t=e.filter((e=>{var t=/(\d+)\.(\d+)\.(\d+)\.(\d+)/g;if(!t.test(e))return e}));return t}},percentageChange(e,t){return 0===e?0:parseInt(e/t*1e4)/100},customColorMethod(e,t){var i=this.percentageChange(e,t);return i<30?"#67c23a":i<70?"#e6a23c":"#f56c6c"},handleSizeChange(e){this.size=e,this.getContact()},handleCurrentChange(e){this.current=e,this.getContact()},addContact(){this.isAdd=!0;const e=this.$refs.forms;e.dialogVisible=!0},subDelete(e){var t=this;t.delLoading=!0,this.$request.fetchDeleteCantact(e).then((e=>{t.$message({message:"删除成功!!!",type:"success"}),t.delLoading=!1,t.getContact()})).catch((e=>{this.$message.error(e.data)}))},getContact(){var e={current:this.current,size:this.size,name:this.selectInput};this.$request.fetchGetContact(e).then((e=>{this.contactData=e.data.records,this.pageTotal=e.data.total}))},handleIgnore(e,t){this.isAddIgnore=!1;const a=this.$refs.formsIgnore;let s=i(8127).Base64;null===t.risk_code&&(t.risk_code=""),null===t.event_code&&(t.event_code=""),a.ignoreShow=!0,a.form={id:t.id,name:t.name,events:t.events,risks:t.risks,ways:t.ways,mobile:t.mobile,dong:t.dong,email:t.email,risk_code:s.decode(t.risk_code),event_code:s.decode(t.event_code)}},handleEdit(e,t){this.isAdd=!1;const i=this.$refs.forms;if(null!==t.risks)var a=t.risks.join("\n");if(null!==t.events)var s=t.events.join("\n");i.dialogVisible=!0,i.form={id:t.id,name:t.name,events:t.events,risks:t.risks,ways:t.ways,mobile:t.mobile,dong:t.dong,email:t.email},i.eventsData=s,i.risksData=a}}},_=y,k=(0,d.A)(_,a,s,!1,null,null,null),w=k.exports}}]);