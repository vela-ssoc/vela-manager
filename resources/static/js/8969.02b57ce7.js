"use strict";(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[8969],{12037:function(e,t,a){a.d(t,{Z:function(){return u}});var r=function(){var e=this,t=e._self._c;return t("div",{class:"事件详情"===e.nameShow?"pageClass":"",staticStyle:{height:"20px","padding-top":"10px"}},[t("span",{staticStyle:{display:"inline-block","font-size":"13px","min-width":"35.5px",height:"28px","line-height":"30px","vertical-align":"top","box-sizing":"border-box",float:"right"}},[e._v("共"+e._s(e.pageTotal)+"条")]),t("el-pagination",{staticStyle:{float:"right"},attrs:{background:"",layout:"jumper,sizes,prev, pager, next","page-size":e.size||15,"current-page":e.pageCurrent,total:e.pageTotal,"page-sizes":e.pageSizeArray||[15,30,45]},on:{"update:currentPage":function(t){e.pageCurrent=t},"update:current-page":function(t){e.pageCurrent=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)},s=[],n={name:"page",props:{pageSizeArray:Array,size:Number,pageTotal:Number,current:Number,nameShow:String},data(){return{pageCurrent:this.current}},methods:{handleSizeChange(e){this.$emit("handleSizeChange",e)},handleCurrentChange(e){this.$emit("handleCurrentChange",e)}},watch:{current:function(e){this.pageCurrent=e}}},i=n,l=a(1001),o=(0,l.Z)(i,r,s,!1,null,"277403d7",null),u=o.exports},98969:function(e,t,a){a.r(t),a.d(t,{default:function(){return b}});var r=function(){var e=this,t=e._self._c;return t("div",[t("el-card",{staticClass:"box-card"},[t("el-row",[t("el-col",{staticStyle:{"text-align":"right"},attrs:{span:24}},[t("el-button",{staticStyle:{"margin-left":"0"},attrs:{"el-button":"",type:"success",plain:"",size:"mini"},on:{click:e.addUser}},[e._v("添加")])],1)],1),t("el-table",{staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:e.usersData,"header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[t("el-table-column",{attrs:{prop:"nickname",label:"昵称"}}),t("el-table-column",{attrs:{prop:"username",label:"用户名"}}),t("el-table-column",{attrs:{prop:"created_at",label:"创建时间"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e.dateTime(t.row.created_at)))]}}])}),t("el-table-column",{attrs:{prop:"updated_at",label:"活动时间"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e.dateTime(t.row.session_at)))]}}])}),t("el-table-column",{attrs:{prop:"enable",label:"状态"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-tag",{attrs:{size:"mini",type:!0===a.row.enable?"success":"danger"}},[e._v(" "+e._s(!0===a.row.enable?"启用":"禁用")+" ")])]}}])}),t("el-table-column",{attrs:{label:"操作"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-popconfirm",{attrs:{title:"确定解绑该账户的OTP授权？解绑后需要重新扫码获取授权"},on:{confirm:function(t){return e.unbindOTP(a.row)}}},[t("el-button",{attrs:{slot:"reference",size:"mini",type:"text"},slot:"reference"},[e._v("OTP解绑 ")])],1),t("el-button",{staticStyle:{"margin-left":"10px"},attrs:{size:"mini",type:"text"},on:{click:function(t){return e.userEdit(a.row)}}},[e._v("编辑 ")]),t("el-popover",{ref:a.row.id,attrs:{placement:"top",width:"200"}},[t("p",[e._v("确定删除?")]),t("div",{staticStyle:{"text-align":"right",margin:"0"}},[t("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(t){e.$refs[a.row.id].doClose()}}},[e._v("取消")]),t("el-button",{attrs:{loading:e.delLoading,type:"primary",size:"mini"},on:{click:function(t){return e.userDelete(a.row.id)}}},[e._v("确定 ")])],1),t("el-button",{staticStyle:{"margin-left":"10px"},attrs:{slot:"reference",type:"text",size:"mini"},slot:"reference"},[e._v("删除")])],1)]}}])})],1),t("Page",{attrs:{size:e.userPage.pageSize,current:e.userPage.current,pageTotal:e.userPage.total},on:{handleSizeChange:e.handleSizeChange,handleCurrentChange:e.handleCurrentChange}})],1),t("user-form",{ref:"userForm",attrs:{userAdd:e.userAdd},on:{getUsers:e.getUsers}}),t("el-col",{attrs:{span:12}})],1)},s=[],n=function(){var e=this,t=e._self._c;return t("el-dialog",{attrs:{"before-close":e.handleClose,"el-dialog":"","close-on-click-modal":!1,visible:e.dialogVisible,title:e.userAdd?"新增":"编辑","append-to-body":"","destroy-on-close":"",width:"500px"},on:{"update:visible":function(t){e.dialogVisible=t}}},[t("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,size:"small","label-width":"90px"}},[t("el-form-item",{attrs:{label:"用户",prop:"username"}},[t("el-input",{staticStyle:{width:"100%"},attrs:{disabled:!e.userAdd},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1),t("el-form-item",{attrs:{label:"昵称",prop:"nickname"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.nickname,callback:function(t){e.$set(e.form,"nickname",t)},expression:"form.nickname"}})],1),1===e.form.domain?t("el-form-item",{attrs:{label:"密码",prop:"password"}},[t("el-input",{staticStyle:{width:"100%"},model:{value:e.form.password,callback:function(t){e.$set(e.form,"password",t)},expression:"form.password"}})],1):e._e(),t("el-form-item",{attrs:{label:"domain",prop:"domain"}},[t("el-select",{staticStyle:{width:"100%"},attrs:{disabled:!e.userAdd},model:{value:e.form.domain,callback:function(t){e.$set(e.form,"domain",t)},expression:"form.domain"}},e._l(e.domainType,(function(e,a){return t("el-option",{key:a,attrs:{value:e.value,label:e.label}})})),1)],1),t("el-form-item",{attrs:{label:"enable",prop:"enable"}},[t("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"rgb(225 221 221)"},model:{value:e.form.enable,callback:function(t){e.$set(e.form,"enable",t)},expression:"form.enable"}})],1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmit("form")}}},[e._v(e._s(e.userAdd?"创建":"修改"))])],1)],1)},i=[],l={name:"userForm",props:{userAdd:Boolean},data(){return{form:{id:"",username:"",nickname:"",password:"",domain:0,enable:!0},dialogVisible:!1,nodeData:[],serviceData:[],codeData:[],rules:{username:[{required:!0,message:"请输入名称",trigger:"blur"}],nickname:[{required:!0,message:"请输入昵称",trigger:"blur"}]},domainType:[{label:"本地账号",value:1},{label:"OA账号",value:2}]}},created(){},methods:{handleClose(){this.dialogVisible=!1,this.form={}},onSubmit(e){this.$refs[e].validate((e=>{e&&(this.userAdd?this.$request.fetchAddUsers(this.form).then((()=>{this.$message({message:"添加成功!!!",type:"success"}),this.dialogVisible=!1,this.form={},this.$emit("getUsers")})).catch((e=>{this.$message.error(e.data)})):this.$request.fetchPatchUsers(this.form).then((()=>{this.$message({message:"修改成功!!!",type:"success"}),this.dialogVisible=!1,this.form={},this.$emit("getUsers")})).catch((e=>{this.$message.error(e.data)})))}))}}},o=l,u=a(1001),d=(0,u.Z)(o,n,i,!1,null,"a12332ea",null),c=d.exports,m=a(12037),p=a(90023),g={name:"index",components:{Page:m.Z,UserForm:c},data(){return{age:0,delLoading:!1,timeShow:!0,userAdd:!1,usersData:[],userPage:{current:1,pageSize:15,total:0}}},mounted(){this.getUsers()},methods:{userEdit(e){this.userAdd=!1;var t=this.$refs.userForm;t.dialogVisible=!0,t.form={id:e.id,username:e.username,nickname:e.nickname,password:e.password,domain:e.domain,enable:e.enable}},dateTime(e){var t=a(30381);return t(e).format("YYYY-MM-DD HH:mm:ss")},userDelete(e){this.$request.fetchDelUser(e).then((()=>{this.delLoading=!1,this.$message({message:"删除成功!!!",type:"success"}),this.getUsers()}))},addUser(){var e=this.$refs.userForm;e.dialogVisible=!0,this.userAdd=!0,e.form={id:null,username:null,nickname:null,password:null,domain:null,enable:null}},getUsers(){this.$request.fetchGetUsers(this.userPage.current,this.userPage.pageSize).then((e=>{this.usersData=e.data.records,this.userPage.total=e.data.total}))},handleSizeChange(e){this.userPage.pageSize=e,this.getUsers()},handleCurrentChange(e){this.userPage.current=e,this.getUsers()},async unbindOTP(e){console.log(e);try{await p.Z.delete("/user/totp",{data:{id:e.id}}),this.$message.success("解除成功")}catch(t){this.$message.error("解除失败")}}}},h=g,f=(0,u.Z)(h,r,s,!1,null,"d0a24a5e",null),b=f.exports}}]);