(window.webpackJsonp=window.webpackJsonp||[]).push([[39],{"/zyf":function(e,t,n){"use strict";var a={name:"page",props:{pageSizeArray:Array,size:Number,pageTotal:Number,current:Number,nameShow:String},data:function(){return{pageCurrent:this.current}},methods:{handleSizeChange:function(e){this.$emit("handleSizeChange",e)},handleCurrentChange:function(e){this.$emit("handleCurrentChange",e)}},watch:{current:function(e){this.pageCurrent=e}}},i=(n("t2r+"),n("KHd+")),r=Object(i.a)(a,(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{class:"事件详情"===e.nameShow?"pageClass":"",staticStyle:{height:"20px","padding-top":"10px"}},[n("span",{staticStyle:{display:"inline-block","font-size":"13px","min-width":"35.5px",height:"28px","line-height":"30px","vertical-align":"top","box-sizing":"border-box",float:"right"}},[e._v("共"+e._s(e.pageTotal)+"条")]),e._v(" "),n("el-pagination",{staticStyle:{float:"right"},attrs:{background:"",layout:"prev, pager, next","page-size":e.size,"current-page":e.pageCurrent,total:e.pageTotal},on:{"update:currentPage":function(t){e.pageCurrent=t},"update:current-page":function(t){e.pageCurrent=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)}),[],!1,null,"40b774d6",null);t.a=r.exports},TzQq:function(e,t,n){},X0WK:function(e,t,n){"use strict";n.r(t);var a={name:"formRelation",props:{isAdd:Boolean},data:function(){return{form:{name:"",tags:[],compounds:[],enable:!0,substances:[],version:"",exclusion:[]},nodeTextarea:"",dialogVisible:!1,tagData:[],serviceData:[],codeData:[],rules:{name:[{required:!0,message:"请输入名称",trigger:"blur"}],tags:[{required:!0,message:"请输入标签",trigger:"change"}]}}},created:function(){this.getNodeList()},methods:{handleClose:function(){this.dialogVisible=!1,this.form={}},getNodeList:function(){var e=this;this.$request.fetchGetTag().then((function(t){void 0!==t&&(e.tagData=t.data)})),this.$request.fetchGetLinks().then((function(t){void 0!==t&&(e.codeData=t.data)})),this.$request.fetchGetServiceList().then((function(t){void 0!==t&&(e.serviceData=t.data)}))},mapArray:function(e){if(null!==e)return e.map((function(e){return e.id}))},onSubmit:function(e){var t=this;void 0!==this.nodeTextarea&&(this.form.exclusion=this.nodeTextarea.split("\n").filter((function(e){return""!==e}))),this.$refs[e].validate((function(e){if(e){var n=t.mapArray(t.form.substances),a=t.mapArray(t.form.compounds),i={name:t.form.name,tags:t.form.tags,compounds:a,substances:n,version:t.form.version,enable:t.form.enable,exclusion:t.form.exclusion};t.isAdd?t.$request.fetchAddRelation(i).then((function(){t.$message({message:"添加成功!!!",type:"success"}),t.dialogVisible=!1,t.form={},i={},t.$parent.getRelation()})).catch((function(e){t.$message.error(e.data),t.form={},i={}})):(i.id=t.form.id,t.$request.fetchPatchRelation(i).then((function(){t.$message({message:"修改成功!!!",type:"success"}),t.dialogVisible=!1,t.form={},t.$parent.getRelation()})).catch((function(e){t.$message.error(e.data),t.form={},i={},t.dialogVisible=!1})))}}))}}},i=n("KHd+"),r={name:"index",components:{RelationForm:Object(i.a)(a,(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-dialog",{attrs:{"before-close":e.handleClose,"el-dialog":"","close-on-click-modal":!1,visible:e.dialogVisible,title:e.isAdd?"新增配置":"编辑配置","append-to-body":"",width:"500px"},on:{"update:visible":function(t){e.dialogVisible=t}}},[n("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,size:"small","label-width":"90px"}},[n("el-form-item",{attrs:{label:"名称",prop:"name"}},[n("el-input",{staticStyle:{width:"100%"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1),e._v(" "),n("el-form-item",{attrs:{label:"标签",prop:"tags"}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"","value-key":"id",filterable:""},model:{value:e.form.tags,callback:function(t){e.$set(e.form,"tags",t)},expression:"form.tags"}},e._l(e.tagData,(function(e,t){return n("el-option",{key:t,attrs:{value:e,label:e}})})),1)],1),e._v(" "),n("el-form-item",{attrs:{label:"服务",prop:"compounds"}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"","value-key":"id"},model:{value:e.form.compounds,callback:function(t){e.$set(e.form,"compounds",t)},expression:"form.compounds"}},e._l(e.serviceData,(function(e){return n("el-option",{key:e.id,attrs:{value:e,label:e.name}})})),1)],1),e._v(" "),n("el-form-item",{attrs:{label:"配置",prop:"substances"}},[n("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"","value-key":"id",filterable:""},model:{value:e.form.substances,callback:function(t){e.$set(e.form,"substances",t)},expression:"form.substances"}},e._l(e.codeData,(function(e){return n("el-option",{key:e.id,attrs:{value:e,label:e.name}})})),1)],1),e._v(" "),n("el-form-item",{attrs:{label:"状态",prop:"enable"}},[n("el-switch",{attrs:{"active-color":"#13ce66","inactive-color":"#ff4949"},model:{value:e.form.enable,callback:function(t){e.$set(e.form,"enable",t)},expression:"form.enable"}})],1),e._v(" "),n("el-form-item",{attrs:{label:"排除节点",prop:"enable"}},[n("el-input",{attrs:{type:"textarea",autosize:{minRows:2,maxRows:4},placeholder:"请输入"},model:{value:e.nodeTextarea,callback:function(t){e.nodeTextarea=t},expression:"nodeTextarea"}})],1)],1),e._v(" "),n("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),e._v(" "),n("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmit("form")}}},[e._v(e._s(e.isAdd?"创建":"修改"))])],1)],1)}),[],!1,null,"7b641561",null).exports,Page:n("/zyf").a},data:function(){return{pageSizeArray:["10","20","30","40"],current:1,size:10,pageTotal:0,relationData:[],isAdd:!0,delLoading:!1,selectInput:""}},created:function(){this.getRelation()},methods:{percentageChange:function(e,t){return 0===e?0:parseInt(e/t*1e4)/100},customColorMethod:function(e,t){var n=this.percentageChange(e,t);return n<30?"#67c23a":n<70?"#e6a23c":"#f56c6c"},enableChange:function(e){var t=this,n={id:e.id,enable:e.enable,version:e.version};this.$request.fetchPatchEnable(n).then((function(){t.getRelation()}))},handleSizeChange:function(e){this.size=e,this.getRelation()},handleCurrentChange:function(e){this.current=e,this.getRelation()},addRelation:function(){this.$refs.forms.dialogVisible=!0},subDelete:function(e){var t=this;this.$refs[e].doClose();var n=this;n.delLoading=!0,this.$request.fetchDelRelation(e).then((function(e){n.$message({message:"删除成功!!!",type:"success"}),n.delLoading=!1,n.getRelation()})).catch((function(e){t.$message.error(e.data)}))},getRelation:function(){var e=this,t={current:this.current,size:this.size,name:this.selectInput};this.$request.fetchGetRelation(t).then((function(t){e.relationData=t.data.records,e.pageTotal=t.data.total}))},handleEdit:function(e,t){this.isAdd=!1;var n=this.$refs.forms;if(n.dialogVisible=!0,null!==t.exclusion)var a=t.exclusion.join("\n");n.form={id:t.id,name:t.name,tags:t.tags,compounds:t.compounds,substances:t.substances,version:t.version,enable:t.enable,exclusion:a},n.nodeTextarea=a},handleDelete:function(e,t){this.$message({showClose:!0,message:e,row:t,type:"success"})},formatter:function(e,t){return e.address},filterTag:function(e,t){return t.tag===e},filterHandler:function(e,t,n){return t[n.property]===e}}},s=Object(i.a)(r,(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticStyle:{width:"100%"}},[n("el-card",[n("el-row",[n("el-col",{attrs:{span:12}},[n("el-input",{staticClass:"input-with-select",staticStyle:{width:"20%"},attrs:{placeholder:"请输入名称",size:"mini"},model:{value:e.selectInput,callback:function(t){e.selectInput=t},expression:"selectInput"}}),e._v(" "),n("el-button",{staticStyle:{color:"#909399",background:"#f4f4f5","border-color":"#d3d4d6"},attrs:{icon:"el-icon-search",size:"mini"},on:{click:e.getRelation}})],1),e._v(" "),n("el-col",{staticStyle:{"text-align":"right"},attrs:{span:12}},[n("el-button",{attrs:{"el-button":"",type:"success",plain:"",size:"mini"},on:{click:e.addRelation}},[e._v("添加")])],1)],1),e._v(" "),n("el-table",{staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:e.relationData,border:"","header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[n("el-table-column",{attrs:{prop:"name",label:"名称",width:"240"}}),e._v(" "),n("el-table-column",{attrs:{label:"标签"},scopedSlots:e._u([{key:"default",fn:function(t){return e._l(t.row.tags,(function(t,a){return n("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v("\n            "+e._s(t)+"\n          ")])}))}}])}),e._v(" "),n("el-table-column",{attrs:{label:"服务"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("div",[n("p",{staticStyle:{float:"left",width:"100%","text-align":"left"}},[e._v("\n              服务：\n              "),e._l(t.row.compounds,(function(t){return n("el-tag",{key:t.id,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v("\n                "+e._s(t.name)+"\n              ")])}))],2),e._v(" "),n("p",{staticStyle:{float:"left",width:"100%","text-align":"left"}},[e._v(" 插件：\n              "),e._l(t.row.substances,(function(t){return n("el-tag",{key:t.id,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v("\n                "+e._s(t.name)+"\n              ")])}))],2)])]}}])}),e._v(" "),n("el-table-column",{attrs:{label:"状态",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-tag",{attrs:{size:"mini",type:!0===t.row.enable?"success":"danger"}},[e._v("\n            "+e._s(!0===t.row.enable?"开":"关")+"\n          ")])]}}])}),e._v(" "),n("el-table-column",{attrs:{prop:"tags",label:"排除节点",width:"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[null!==t.row.exclusion?n("span",[t.row.exclusion.length<=1?n("span",e._l(t.row.exclusion,(function(t,a){return n("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v("\n                      "+e._s(t))])})),1):n("span",[e._l(t.row.exclusion.slice(0,1),(function(t,a){return n("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v(" "+e._s(t))])})),e._v(" "),n("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"Right Center 提示文字",placement:"right"}},[n("el-tag",{attrs:{size:"mini",type:"info"}},[e._v(" + "+e._s(t.row.exclusion.length))]),e._v(" "),n("div",{attrs:{slot:"content"},slot:"content"},e._l(t.row.exclusion,(function(t,a){return n("span",{key:a,staticStyle:{"padding-right":"15px"}},[e._v(e._s(t)),n("br")])})),0)],1)],2)]):e._e()]}}])}),e._v(" "),n("el-table-column",{attrs:{label:"操作",width:"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(n){return e.handleEdit(t.$index,t.row)}}},[e._v("编辑\n          ")]),e._v(" "),n("el-popover",{ref:t.row.id,attrs:{placement:"top"}},[n("p",[e._v("确定删除?")]),e._v(" "),n("div",{staticStyle:{"text-align":"right",margin:"0"}},[n("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(n){e.$refs[t.row.id].doClose()}}},[e._v("取消")]),e._v(" "),n("el-button",{attrs:{loading:e.delLoading,type:"primary",size:"mini"},on:{click:function(n){return e.subDelete(t.row.id)}}},[e._v("确定\n              ")])],1),e._v(" "),n("el-button",{attrs:{slot:"reference",type:"text",size:"mini"},slot:"reference"},[e._v("删除")])],1)]}}])})],1),e._v(" "),n("Page",{attrs:{size:e.size,current:e.current,pageTotal:e.pageTotal,pageSizeArray:e.pageSizeArray},on:{handleSizeChange:e.handleSizeChange,handleCurrentChange:e.handleCurrentChange}})],1),e._v(" "),n("RelationForm",{ref:"forms",attrs:{"is-add":e.isAdd}})],1)}),[],!1,null,null,null);t.default=s.exports},"t2r+":function(e,t,n){"use strict";n("TzQq")}}]);
//# sourceMappingURL=39.8580908943b10c476969.js.map