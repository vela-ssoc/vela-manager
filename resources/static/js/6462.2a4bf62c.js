"use strict";(self["webpackChunkssoc"]=self["webpackChunkssoc"]||[]).push([[6462],{72953:function(e,t,i){i.d(t,{A:function(){return y}});var a=function(){var e=this,t=e._self._c;return t("div",{staticStyle:{display:"flex","flex-wrap":"wrap","align-items":"center"}},[e._t("prefix"),e._l(e.tabs,(function(i,a){return t("el-popover",{key:a,attrs:{"popper-class":"filterOptionsPopover",placement:e.placement,width:"280",trigger:"click"},on:{show:function(t){return e.popovershow(a)}},scopedSlots:e._u([{key:"reference",fn:function(){return[t("span",[0===a?t("el-button",{staticClass:"buttonsearch",staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._t("textSlot",(function(){return[e._v("添加条件")]}))],2):[t("span",{key:a,staticStyle:{margin:"2px","font-size":"12px",color:"#fff",padding:"3px","border-radius":"3px","background-color":"#409eff",display:"inline-block"},on:{click:function(t){return e.editTag(t,i,a-1)}}},[e._v(" "+e._s(i.keyDesc)+" "+e._s(i.operatorDesc)+" "+e._s(i.valueDesc)+" "),i.clearable?t("i",{staticClass:"el-icon-close",staticStyle:{cursor:"pointer"},on:{click:function(t){return t.stopPropagation(),e.delTag(i,a-1)}}}):e._e()])]],2)]},proxy:!0}],null,!0),model:{value:i.show,callback:function(t){e.$set(i,"show",t)},expression:"item.show"}},[t("el-form",{ref:"formRef",refInFor:!0,attrs:{model:e.form,"label-position":"right","label-width":"70px"}},[t("el-form-item",{attrs:{label:"关键字",prop:"key",rules:e.required}},[t("el-select",{staticStyle:{width:"100%"},attrs:{filterable:"",size:"mini"},on:{change:e.selectField},model:{value:e.form.key,callback:function(t){e.$set(e.form,"key",t)},expression:"form.key"}},e._l(e.options,(function(e,i){return t("el-option",{key:i,attrs:{value:e.key,label:e.desc}})})),1)],1),t("el-form-item",{attrs:{label:"操作符",prop:"operator",rules:e.required}},[t("el-select",{staticStyle:{width:"100%"},attrs:{size:"mini"},model:{value:e.form.operator,callback:function(t){e.$set(e.form,"operator",t)},expression:"form.operator"}},e._l(e.selectFieldDerive.operatorArray,(function(e,i){return t("el-option",{key:i,attrs:{value:e.key,label:e.desc}})})),1)],1),t("el-form-item",{attrs:{label:"值",rules:e.required,prop:"value"}},[t("div",{directives:[{name:"show",rawName:"v-show",value:["datetime","time"].includes(e.selectFieldDerive.valueType),expression:"['datetime', 'time'].includes(selectFieldDerive.valueType)"}],on:{click:function(e){e.stopPropagation()}}},[t("el-date-picker",{ref:"dateRef",refInFor:!0,staticStyle:{width:"100%"},attrs:{clearable:"","popper-class":"filterDate",teleported:!1,"append-to-body":!1,type:"datetime",pickerOptions:e.pickerOptions,placeholder:"选择时间",size:"mini"},model:{value:e.form.value,callback:function(t){e.$set(e.form,"value",t)},expression:"form.value"}})],1),["datetime","time"].includes(e.selectFieldDerive.valueType)?e._e():[e.selectFieldDerive.isEnum?t("el-select",{staticStyle:{width:"100%"},attrs:{size:"mini",clearable:"",multiple:"in"===e.form.operator||"notin"===e.form.operator,"collapse-tags":""},model:{value:e.form.value,callback:function(t){e.$set(e.form,"value",t)},expression:"form.value"}},e._l(e.selectFieldDerive.valueArray,(function(e,i){return t("el-option",{key:i,attrs:{value:e.key,label:e.desc}})})),1):[e.suggestOptions&&e.suggestOptions[e.form.key]?t("el-autocomplete",{staticStyle:{width:"100%"},attrs:{size:"mini","fetch-suggestions":e.querySearch,clearable:""},model:{value:e.form.value,callback:function(t){e.$set(e.form,"value",t)},expression:"form.value"}}):t("el-input",{staticStyle:{width:"100%"},attrs:{size:"mini",clearable:""},model:{value:e.form.value,callback:function(t){e.$set(e.form,"value",t)},expression:"form.value"}})]]],2),t("el-form-item",[t("el-button",{staticClass:"buttonsearch",attrs:{size:"mini"},on:{click:function(t){return e.confirm(a)}}},[e._v("确定")])],1)],1)],1)}))],2)},l=[],s=i(95093),o=i.n(s);class n{constructor(e){this.key=e?.key,this.operator=e?.operator,this.value=e?.value}}const r=o()().startOf("days"),c=o()().startOf("days").subtract(1,"days"),u=o()().startOf("days").subtract(7,"days"),d=o()().startOf("days").subtract(15,"days"),h=o()().startOf("days").subtract(30,"days"),p={[r]:"今天",[c]:"昨天",[u]:"七天前",[d]:"十五天前",[h]:"三十天前"};var f={props:{value:{type:Array,default:()=>[]},options:{default:()=>[]},placement:{default:"top"},initData:{type:Object,default:null},suggestOptions:{default:()=>({})},multiCondition:{default:!0}},created(){},mounted(){},data(){return{moment:Object.freeze(o()),form:new n,show:!1,editIdx:-1,pickerOptions:{disabledDate(e){return e.getTime()>Date.now()},shortcuts:[{text:"今天",onClick(e){e.$emit("pick",r)}},{text:"昨天",onClick(e){e.$emit("pick",c)}},{text:"七天前",onClick(e){e.$emit("pick",u)}},{text:"十五天前",onClick(e){e.$emit("pick",d)}},{text:"三十天前",onClick(e){e.$emit("pick",h)}}]},required:{required:!0,message:"请检查必填项",trigger:["blur","change"]}}},computed:{filters:{get(){return this.value?.map((e=>this.getSelectItem(e)))},set(e){this.$emit("input",e),this.$emit("change",e)}},selectFieldDerive(){var e=this.options.findIndex((e=>e.key===this.form.key));return{operatorArray:this.options?.[e]?.operators||[],isEnum:this.options?.[e]?.enum,valueArray:this.options?.[e]?.enums||[],valueType:this.options?.[e]?.type}},tabs(){return[{show:!1},...this.filters]}},methods:{querySearch(e,t){var i=this.suggestOptions?.[this.form.key],a=e?i.filter((t=>-1!==String(t).toLowerCase().indexOf(e.toLowerCase()))):i;t(a.map((e=>({value:e}))))},selectField(e){this.form.value=null,this.$nextTick((()=>{this.form.operator=this.selectFieldDerive.operatorArray.some((e=>"like"===e.key))?"like":this.selectFieldDerive.operatorArray?.[0]?.key}))},selectOperator(e){this.form.value=null},async confirm(e){try{if(await(this.$refs.formRef?.[e].validate().catch((()=>{throw"请检查必填项"}))),this.show=!1,this.editIdx>=0&&this.editIdx<this.filters.length){const e=this.filters||[];e[this.editIdx]=this.valueDecode(Object.assign(e[this.editIdx],this.form)),this.filters=[...e],this.editIdx=-1}else this.pushFilter(this.form)}catch(t){this.$message.warning(t)}},valueDecode(e){const t=this.options?.find((t=>t.key===e.key));return t.enum&&e.value&&("in"===e.operator||"notin"===e.operator)&&Array.isArray(e.value)?{...e,value:e.value.join(",")}:{...e}},valueEncode(e){const t=this.options?.find((t=>t.key===e.key));return t?.enum&&e.value&&("in"===e.operator||"notin"===e.operator)&&(e.value=e.value?.split(",")||e.value),e},pushFilter(e){const t=this.filters||[],i=t.findIndex((t=>t.key===e.key)),a=t[i];e=this.valueDecode(e),a?(null===e.value||void 0===e||""===e?t.splice(i,1):Object.assign(a,e),this.filters=[...t]):null!==e.value&&void 0!==e&&""!==e&&(!this.multiCondition&&t.length>0&&(t.length=0),this.filters=[...t,{...e}])},getSelectItem({key:e,operator:t,value:i,clearable:a=!0,disabled:l=!1}){const s=this.options?.find((t=>t.key===e)),n=s?.operators,r=s?.enums||[],c=()=>{const a=this.valueEncode({key:e,operator:t,value:i})?.value;if("time"===s?.type){const e=o()(i);return e in p?p[e]:e.format("YYYY-MM-DD HH:mm:ss")}return s?.enum?"in"===t||"notin"===t?a?.map((e=>r.find((t=>t.key===e))?.desc))?.join(",")||a:r.find((e=>e.key===a))?.desc:i};return{key:e,keyDesc:s?.desc,value:i,valueDesc:c(),show:!1,valueType:s?.type,operator:t,clearable:a,disabled:l,operatorDesc:n?.find((e=>e.key===t))?.desc}},delTag(e,t){const i=this.filters||[];i.splice(t,1),this.filters=i},editTag(e,t,i){t.disabled?e.stopPropagation():(this.editIdx=i,this.form=new n(this.valueEncode(t)))},popovershow(e){this.$refs.dateRef[e].placement="right"}}},m=f,g=i(81656),b=(0,g.A)(m,a,l,!1,null,"980cc6ea",null),y=b.exports},92161:function(e,t,i){i.d(t,{A:function(){return c}});var a=function(){var e=this,t=e._self._c;return t("div",{class:"事件详情"===e.nameShow?"pageClass":"",staticStyle:{height:"20px","padding-top":"10px"}},[t("span",{staticStyle:{display:"inline-block","font-size":"13px","min-width":"35.5px",height:"28px","line-height":"30px","vertical-align":"top","box-sizing":"border-box",float:"right"}},[e._v("共"+e._s(e.pageTotal)+"条")]),t("el-pagination",{staticStyle:{float:"right"},attrs:{background:"",layout:"jumper,sizes,prev, pager, next","page-size":e.size||15,"current-page":e.pageCurrent,total:e.pageTotal,"page-sizes":e.pageSizeArray||[15,30,45]},on:{"update:currentPage":function(t){e.pageCurrent=t},"update:current-page":function(t){e.pageCurrent=t},"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1)},l=[],s={name:"page",props:{pageSizeArray:Array,size:Number,pageTotal:Number,current:Number,nameShow:String},data(){return{pageCurrent:this.current}},methods:{handleSizeChange(e){this.$emit("handleSizeChange",e)},handleCurrentChange(e){this.$emit("handleCurrentChange",e)}},watch:{current:function(e){this.pageCurrent=e}}},o=s,n=i(81656),r=(0,n.A)(o,a,l,!1,null,"277403d7",null),c=r.exports},32268:function(e,t,i){i.d(t,{A:function(){return x}});var a=function(){var e=this,t=e._self._c;return t("div",{staticStyle:{width:"100%"}},[t("div",{staticClass:"tableHead"},[t("div",{staticStyle:{flex:"1","margin-right":"5px"}},[e._t("right_top_btn")],2),t("span",{staticClass:"icon"},[e.refreshAble?t("i",{staticClass:"el-icon-refresh",attrs:{title:"重加载数据"},on:{click:e.loadData}}):e._e(),e.configAble&&"inhead"!==e.configIconPos?t("i",{class:{"el-icon-setting":!0},attrs:{title:"编辑表格"},on:{click:function(t){e.dialogVisible=!0}}}):e._e()])]),t("div",{staticStyle:{position:"relative"}},[e.configAble&&"inhead"===e.configIconPos?t("i",{staticClass:"el-icon-setting iconInHead",attrs:{title:"编辑表格"},on:{click:function(t){e.dialogVisible=!0}}}):e._e()]),t("el-table",e._g(e._b({ref:"table",staticStyle:{width:"100%"},attrs:{"header-cell-style":e.$attrs["header-cell-style"]||{color:"#909399",textAlign:"center",background:"#f5f7fa"},data:e.showData,stripe:e.stripe},on:{"selection-change":e.selectionChange}},"el-table",e.$attrs,!1),e.$listeners),[e.selection?t("el-table-column",{attrs:{selectable:e.selectable,type:"selection"}}):e._e(),e.index?t("el-table-column",{attrs:{type:"index",label:"序号",align:"center",fixed:"left"}}):e._e(),e._l(e.showColumn,(function(i,a){return[t("el-table-column",{key:a,attrs:{prop:i&&i.prop,label:i&&i.label,width:i&&i.width,"min-width":i&&i["min-width"],fixed:i&&i.fixed,formatter:i&&i.formatter,"show-overflow-tooltip":i&&i["show-overflow-tooltip"]||!1,sortable:i&&i.sortable,"sort-method":i&&i["sort-method"],"sort-by":i&&i["sort-by"],align:i&&i.align||"center","sort-orders":i&&i["sort-orders"],type:i&&i.type,"class-name":i&&i["class-name"]},scopedSlots:e._u([i&&i.slotHeader?{key:"header",fn:function(t){return[e._t(i&&i.slotHeader,null,{row:t.row,col:i,$index:t.$index})]}}:null,i&&i.slot?{key:"default",fn:function(t){return[e._t(i&&i.slot,null,{row:t.row,col:i,$index:t.$index})]}}:null],null,!0)})]}))],2),e.needPagination?t("div",{style:e.paginationStyle},[t("el-pagination",{attrs:{"current-page":e.pagination.pageIndex,"page-sizes":e.pagination.pageSizes||[15,30,50],"page-size":e.pagination.pageSize,layout:e.pagination.layout||"jumper, sizes, prev, pager, next,total",total:e.pagination.total},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange,"update:currentPage":function(t){return e.$set(e.pagination,"pageIndex",t)},"update:current-page":function(t){return e.$set(e.pagination,"pageIndex",t)},"update:pageSize":function(t){return e.$set(e.pagination,"pageSize",t)},"update:page-size":function(t){return e.$set(e.pagination,"pageSize",t)}}})],1):e._e(),t("ColumnConfig",{attrs:{columns:e.filterColumns,visible:e.dialogVisible,selectColumn:e.selectedColumn},on:{"update:visible":function(t){e.dialogVisible=t},"update:selectColumn":function(t){e.selectedColumn=t},"update:select-column":function(t){e.selectedColumn=t}}})],1)},l=[],s=function(){var e=this,t=e._self._c;return t("div",[t("el-dialog",{attrs:{title:"设置需展示的列",visible:e.dialogVisible,width:"800px","append-to-body":""},on:{open:function(t){return e.opened()},"update:visible":function(t){e.dialogVisible=t}}},[t("div",[t("el-checkbox",{attrs:{indeterminate:e.isIndeterminate},on:{change:e.handleCheckAllChange},model:{value:e.checkAll,callback:function(t){e.checkAll=t},expression:"checkAll"}},[e._v("全选")]),t("el-button",{staticStyle:{"text-decoration":"underline",color:"red"},attrs:{type:"text"},on:{click:function(t){return e.$refs.SortColumn.open()}}},[e._v("手动排序")]),t("el-button",{staticStyle:{"text-decoration":"underline",color:"red"},attrs:{type:"text"},on:{click:e.reset}},[e._v("重置配置")]),t("div",{staticStyle:{margin:"15px 0"}}),t("el-checkbox-group",{on:{change:e.handleCheckedCitiesChange},model:{value:e.checkboxSelectBuff,callback:function(t){e.checkboxSelectBuff=t},expression:"checkboxSelectBuff"}},e._l(e.__column,(function(i){return t("el-col",{key:i.columnUUkey,staticClass:"colCheckbox",attrs:{span:8}},[t("el-checkbox",{attrs:{title:i.label,label:i.columnUUkey}},[e._v(e._s(i.label)+" ")])],1)})),1)],1),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{attrs:{size:"small"},on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{size:"small",type:"primary"},on:{click:e.confirm}},[e._v("确 定")])],1)]),t("SortColumn",{ref:"SortColumn",attrs:{column:e.__column},on:{close:e.clsoeSort}})],1)},o=[],n=function(){var e=this,t=e._self._c;return t("el-dialog",{attrs:{title:"拖动节点进行排序",visible:e.dialogVisible,width:"500px","append-to-body":""},on:{"update:visible":function(t){e.dialogVisible=t}}},[t("el-tree",{attrs:{data:e.data,"node-key":"columnUUkey","default-expand-all":"",draggable:"","allow-drop":e.allowDrog}}),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{attrs:{size:"small",type:"primary"},on:{click:e.close}},[e._v("确 定")])],1)],1)},r=[],c={props:{column:{type:Array,require:!0}},data(){return{data:[],dialogVisible:!1}},created(){},methods:{open(){this.dialogVisible=!0,this.data=this.column.map((e=>({columnUUkey:e.columnUUkey,label:e.label})))},close(){this.dialogVisible=!1,this.$emit("close",this.data.map((e=>e.columnUUkey)))},allowDrog(e,t,i){return"inner"!==i}}},u=c,d=i(81656),h=(0,d.A)(u,n,r,!1,null,null,null),p=h.exports,f={props:{columns:{type:Array,default:()=>[]},visible:{type:Boolean,require:!0,default:!1},selectColumn:{type:Array,default:()=>[]}},components:{SortColumn:p},computed:{dialogVisible:{get(){return this.visible},set(e){this.$emit("update:visible",e)}},checkboxSelect:{get(){return this.selectColumn},set(e){this.$emit("update:selectColumn",e)}},__column(){return this.columns.map((e=>({label:e.label,columnUUkey:e.prop||e.slot}))).sort(((e,t)=>(this.__sortWeight[e.columnUUkey]||0)-(this.__sortWeight[t.columnUUkey]||0)))},__sortWeight(){const e={};for(let t=0;t<this.sortColumnList.length;t++)e[this.sortColumnList[t]]=t;return e}},data(){return{sortColumnList:[],checkAll:!1,isIndeterminate:!1,checkboxSelectBuff:[],oldColumns:[]}},methods:{handleCheckAllChange(e){this.checkboxSelectBuff=e?this.__column.map((e=>e.columnUUkey)):[],this.isIndeterminate=!1},handleCheckedCitiesChange(e){this.checkboxSelectBuff=e;let t=e.length;this.checkAll=t===this.__column.length,this.isIndeterminate=t>0&&t<this.__column.length},opened(){const e=this.__column.filter((e=>!this.oldColumns.includes(e.columnUUkey))).map((e=>e.columnUUkey));this.checkboxSelectBuff=this.checkboxSelectBuff.concat(e).sort(((e,t)=>(this.__sortWeight[e]||0)-(this.__sortWeight[t]||0))),this.oldColumns=this.__column.map((e=>e.columnUUkey)),this.checkAll=this.checkboxSelectBuff.length===this.__column.length,this.isIndeterminate=this.checkboxSelectBuff.length>0&&this.checkboxSelectBuff.length<this.__column.length},clsoeSort(e){this.sortColumnList=this.oldColumns=e||[],this.checkboxSelectBuff=this.checkboxSelectBuff.sort(((e,t)=>(this.__sortWeight[e]||0)-(this.__sortWeight[t]||0)))},reset(){this.sortColumnList=[],this.checkboxSelectBuff=this.__column.map((e=>e.columnUUkey))},confirm(){this.checkboxSelect=this.checkboxSelectBuff,this.sortColumnList=this.oldColumns,this.dialogVisible=!1}}},m=f,g=(0,d.A)(m,s,o,!1,null,"25f6eb76",null),b=g.exports,y={name:"table_list",components:{ColumnConfig:b},props:{data:{type:Array,required:!0,default:()=>[]},pagination:{type:Object,default:()=>({pageIndex:1,pageSize:15,total:0,pageSizes:[15,30,50],layout:"jumper, sizes, prev, pager, next,total"})},localPaginate:{type:Boolean,default:!1},needPagination:{type:Boolean,default:!0},columns:{required:!0,type:Array,default:()=>[]},index:{type:Boolean,default:!1},selection:{type:Boolean,default:!1},selectable:{type:Function,default:()=>()=>{}},selectCountLimit:{type:Number,default:1e6},checkedData:{type:Array,default:()=>[]},configAble:{type:[Boolean],default:!0},configIconPos:{default:"",type:String},refreshAble:{type:Boolean,default:!1},stripe:{type:Boolean,default:!1}},computed:{paginationStyle(){return{display:"flex","padding-top":"10px","flex-direction":"row","justify-content":"left"===this.pagination.align?"flex-start":"flex-end"}},filterColumns(){return this.columns.filter((e=>e.hidden&&"function"===typeof e.hidden?!e.hidden():!e.hidden))},showColumn(){let e=this.filterColumns;return this.selectedColumn.length&&(e=this.selectedColumn.map((e=>this.filterColumns.find((t=>t.prop===e||t.slot===e)))).filter((e=>!!e))),e.map((e=>({...e,columnUUkey:e.prop||e.slot})))},showData(){if(!this.data||!this.data.length)return[];if(!this.needPagination)return this.data;const e=Math.max(this.pagination.pageSize*(this.pagination.pageIndex-1),0),t=Math.min(this.data.length,e+this.pagination.pageSize);return this.data.length===this.pagination.total?this.data.slice(e,t):this.data},_checkedData:{get(){return this.checkedData},set(e){this.$emit("update:checkedData",e)}}},data(){return{selectedColumn:[],dialogVisible:!1}},watch:{data:{handler(e,t){this.resetSelect(),this.localPaginate&&(this.pagination.total=this.data&&this.data.length||0)},immediate:!0},filterColumns(){this.selectedColumn=[]}},mounted(){this.resetSelect()},methods:{loadData(){this.$emit("loadData")},reload(){this.pagination.pageIndex=1,this._checkedData=[],this.$emit("loadData")},selectionChange(e,t){if(e&&e.length>this.selectCountLimit){this.$refs.table.clearSelection();for(let t=0;t<this.selectCountLimit;t++)this.$refs.table.toggleRowSelection(e[e.length-t-1],!0)}else this._checkedData=e},resetSelect(){if(this.$refs.table){this.$refs.table.clearSelection();for(let e of this.data)for(let t of this._checkedData)JSON.stringify(e)===JSON.stringify(t)&&this.$refs.table.toggleRowSelection(e,!0)}},handleSizeChange(){this.localPaginate||(this.pagination.pageIndex=1,this.loadData())},handleCurrentChange(){!this.localPaginate&&this.loadData()}}},v=y,k=(0,d.A)(v,a,l,!1,null,"6202e7cb",null),x=k.exports},6462:function(e,t,i){i.r(t),i.d(t,{default:function(){return V}});var a=function(){var e=this,t=e._self._c;return t("div",{staticStyle:{width:"100%"}},[t("el-card",{staticClass:"contentCard_"},[t("el-row",[t("el-col",{attrs:{span:12}},[t("el-input",{staticClass:"input-with-select",staticStyle:{width:"20%"},attrs:{clearable:"",placeholder:"请输入名称",size:"mini"},on:{blur:e.getRelation},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.getRelation.apply(null,arguments)}},model:{value:e.selectInput,callback:function(t){e.selectInput=t},expression:"selectInput"}}),t("el-button",{staticStyle:{color:"#909399",background:"#f4f4f5","border-color":"#d3d4d6"},attrs:{icon:"el-icon-search",size:"mini"},on:{click:e.getRelation}})],1),t("el-col",{staticStyle:{"text-align":"right"},attrs:{span:12}},[t("el-button",{attrs:{type:"success",plain:"",size:"mini"},on:{click:e.addRelation}},[e._v("添加")]),t("el-button",{attrs:{type:"success",plain:"",size:"mini"},on:{click:function(t){return e.$refs.progressDetailRef.openProgress()}}},[e._v("任务进度")])],1)],1),t("el-table",{staticStyle:{width:"100%","margin-top":"5px"},attrs:{data:e.relationData,"header-cell-style":{color:"#909399",textAlign:"center",background:"#f5f7fa"}}},[t("el-table-column",{attrs:{prop:"name",label:"名称",width:"240"}}),t("el-table-column",{attrs:{label:"标签"},scopedSlots:e._u([{key:"default",fn:function(i){return[e._l(i.row.tags&&i.row.tags.slice(0,2)||[],(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v(" "+e._s(i)+" ")])})),i.row.tags&&i.row.tags.length>2?t("el-popover",{attrs:{placement:"top",width:"200",trigger:"hover"}},[t("span",e._l(i.row.tags,(function(i,a){return t("el-tag",{key:a,staticStyle:{margin:"3px"},attrs:{size:"mini"}},[e._v(" "+e._s(i)+" ")])})),1),t("el-tag",{attrs:{slot:"reference",size:"mini",type:"info"},slot:"reference"},[e._v(" + "+e._s(i.row.tags.length))])],1):e._e()]}}])}),t("el-table-column",{attrs:{label:"服务"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("div",[t("p",{staticStyle:{float:"left",width:"100%","text-align":"left"}},[e._v(" 插件： "),e._l(i.row.substances,(function(i){return t("el-tag",{key:i.id,staticStyle:{"margin-right":"3px"},attrs:{type:"info",size:"mini"}},[e._v(" "+e._s(i.name)+" ")])}))],2)])]}}])}),t("el-table-column",{attrs:{label:"状态",width:"150"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-tag",{attrs:{size:"mini",type:!0===i.row.enable?"success":"danger"}},[e._v(" "+e._s(!0===i.row.enable?"开":"关")+" ")])]}}])}),t("el-table-column",{attrs:{prop:"tags",label:"排除节点",width:"180"},scopedSlots:e._u([{key:"default",fn:function(i){return[null!==i.row.exclusion?t("span",[i.row.exclusion.length<=1?t("span",e._l(i.row.exclusion,(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v(" "+e._s(i))])})),1):t("span",[e._l(i.row.exclusion.slice(0,1),(function(i,a){return t("el-tag",{key:a,staticStyle:{"margin-right":"3px"},attrs:{size:"mini"}},[e._v(" "+e._s(i))])})),t("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:"Right Center 提示文字",placement:"right"}},[t("el-tag",{attrs:{size:"mini",type:"info"}},[e._v(" + "+e._s(i.row.exclusion.length))]),t("div",{attrs:{slot:"content"},slot:"content"},e._l(i.row.exclusion,(function(i,a){return t("span",{key:a,staticStyle:{"padding-right":"15px"}},[e._v(e._s(i)),t("br")])})),0)],1)],2)]):e._e()]}}])}),t("el-table-column",{attrs:{label:"操作",width:"180"},scopedSlots:e._u([{key:"default",fn:function(i){return[t("el-button",{staticStyle:{"margin-right":"5px"},attrs:{disabled:100!==e.percent,size:"mini",type:"text"},on:{click:function(t){return e.handleEdit(i.$index,i.row)}}},[e._v("编辑 ")]),t("el-popover",{ref:i.row.id,attrs:{placement:"top"}},[t("p",[e._v("确定删除?")]),t("div",{staticStyle:{"text-align":"right",margin:"0"}},[t("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(t){e.$refs[i.row.id].doClose()}}},[e._v("取消")]),t("el-button",{attrs:{loading:e.delLoading,type:"primary",size:"mini"},on:{click:function(t){return e.subDelete(i.row.id)}}},[e._v("确定 ")])],1),t("el-button",{staticStyle:{color:"red"},attrs:{slot:"reference",type:"text",size:"mini"},slot:"reference"},[e._v("删除")])],1)]}}])})],1),t("Page",{attrs:{size:e.size,current:e.current,pageTotal:e.pageTotal},on:{handleSizeChange:e.handleSizeChange,handleCurrentChange:e.handleCurrentChange}})],1),t("RelationForm",{ref:"forms",attrs:{"is-add":e.isAdd},on:{loadData:()=>{e.showProgress=!0,e.getRelation()}}}),t("progressDetail",{ref:"progressDetailRef"})],1)},l=[],s=function(){var e=this,t=e._self._c;return t("el-drawer",{attrs:{"before-close":e.handleClose,visible:e.dialogVisible,"close-on-click-modal":!1,wrapperClosable:!1,title:e.isAdd?"新增配置":"编辑配置",size:"500px"},on:{"update:visible":function(t){e.dialogVisible=t}}},[t("div",{staticStyle:{"border-top":"1px solid #ebe1e1",height:"100%",display:"flex","flex-direction":"column",padding:"10px"}},[t("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,size:"small","label-width":"auto"}},[t("el-form-item",{attrs:{label:"名称",prop:"name"}},[t("el-input",{model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1),t("el-form-item",{attrs:{label:"标签",prop:"tags"}},[t("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"","reserve-keyword":"",clearable:"",placeholder:"可输入关键字远程筛选","remote-method":e.lodash.debounce((t=>{e.remoteMethod(t)}),300),remote:"",filterable:""},model:{value:e.form.tags,callback:function(t){e.$set(e.form,"tags",t)},expression:"form.tags"}},e._l(e.tagData,(function(e){return t("el-option",{key:e,attrs:{value:e,label:e}})})),1)],1),t("el-form-item",{attrs:{label:"配置",prop:"substances"}},[t("el-select",{staticStyle:{width:"100%"},attrs:{multiple:"","reserve-keyword":"",clearable:"",placeholder:"可输入关键字远程筛选",filterable:""},model:{value:e.form.substances,callback:function(t){e.$set(e.form,"substances",t)},expression:"form.substances"}},e._l(e.codeData?.map((e=>({label:e.name,value:e.id}))),(function(e){return t("el-option",{key:e.value,attrs:{value:e.value,label:e.label}})})),1)],1),t("el-form-item",{attrs:{label:"排除节点",prop:"enable"}},[t("el-input",{attrs:{type:"textarea",autosize:{minRows:4},placeholder:"请输入"},model:{value:e.nodeTextarea,callback:function(t){e.nodeTextarea=t},expression:"nodeTextarea"}})],1),t("el-form-item",{attrs:{label:"状态",prop:"enable"}},[t("el-switch",{attrs:{"active-text":"开","inactive-text":"关","active-color":"#13ce66","inactive-color":"#d9d9d9"},model:{value:e.form.enable,callback:function(t){e.$set(e.form,"enable",t)},expression:"form.enable"}})],1)],1),t("div",{staticClass:"dialog-footer",staticStyle:{display:"flex","justify-content":"flex-end","margin-bottom":"0"}},[t("el-button",{attrs:{size:"small"},on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),t("el-button",{attrs:{size:"small",type:"primary"},on:{click:function(t){return e.onSubmit("form")}}},[e._v(e._s(e.isAdd?"创建":"修改"))])],1)],1)])},o=[],n=i(2543),r=i.n(n),c=i(72505),u=i.n(c);function d(e,t,i){return(t=h(t))in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}function h(e){var t=p(e,"string");return"symbol"==typeof t?t:t+""}function p(e,t){if("object"!=typeof e||!e)return e;var i=e[Symbol.toPrimitive];if(void 0!==i){var a=i.call(e,t||"default");if("object"!=typeof a)return a;throw new TypeError("@@toPrimitive must return a primitive value.")}return("string"===t?String:Number)(e)}class f{constructor(){d(this,"name",""),d(this,"tags",[]),d(this,"compounds",[]),d(this,"enable",!0),d(this,"substances",[]),d(this,"version",""),d(this,"exclusion",[])}}var m={name:"formRelation",props:{isAdd:Boolean},components:{},data(){return{something:[],lodash:r(),form:new f,nodeTextarea:"",dialogVisible:!1,tagData:[],codeData:[],rules:{name:[{required:!0,message:"请输入名称",trigger:"blur"}],tags:[{required:!0,message:"请输入标签",trigger:"change"}]},timer:null}},created(){this.getNodeList()},methods:{filterMethod(e,t){return t.name?.includes(e)},remoteMethod(e){this.$request.fetchGetTag(e).then((e=>{void 0!==e&&(this.tagData=e.data?.slice(0,10))}))},handleClose(){this.dialogVisible=!1,this.form={}},getNodeList(){this.$request.fetchGetTag().then((e=>{void 0!==e&&(this.tagData=e.data)})),this.$request.fetchGetLinks().then((e=>{void 0!==e&&(this.codeData=e.data)}))},mapArray(e){if(null!==e)return e.map((e=>e.id))},onSubmit(e){void 0!==this.nodeTextarea&&(this.form.exclusion=this.nodeTextarea.split("\n").filter((e=>""!==e))),this.$refs[e].validate((async e=>{try{if(!e)throw"检查填写项";let i=this.mapArray(this.form.compounds);var t={name:this.form.name,tags:this.form.tags,compounds:i,substances:this.form.substances,version:this.form.version,enable:this.form.enable,exclusion:this.form.exclusion,id:this.form.id};let a=this.isAdd?this.$request.fetchAddRelation:this.$request.fetchPatchRelation;await a(t),this.$message({message:"添加成功!!!",type:"success"}),this.$emit("loadData"),this.form=new f,this.dialogVisible=!1}catch(i){console.error(i),this.$message.error(i)}}))},openDialog(){this.dialogVisible=!0,this.form=new f}},beforeDestroy(){clearTimeout(this.timer)}},g=m,b=i(81656),y=(0,b.A)(g,s,o,!1,null,"17b97912",null),v=y.exports,k=i(92161),x=function(){var e=this,t=e._self._c;return t("el-drawer",{attrs:{withHeader:!1,visible:e.visible,direction:"rtl","destroy-on-close":"",size:"60%"},on:{"update:visible":function(t){e.visible=t},opened:e.domReady,closed:e.close}},[t("div",[t("div",{ref:"echartRef",staticStyle:{width:"100%",height:"300px"}})]),t("div",{staticStyle:{padding:"10px"}},[t("tableList",{ref:"table",attrs:{data:e.tableData,pagination:e.pagination,refreshAble:"",columns:e.columns},on:{loadData:e.loadData},scopedSlots:e._u([{key:"right_top_btn",fn:function(){return[t("div",{staticStyle:{display:"flex"}},[t("el-radio-group",{attrs:{size:"mini"},on:{change:function(t){return e.$refs.table.reload()}},model:{value:e.taskType,callback:function(t){e.taskType=t},expression:"taskType"}},[t("el-radio-button",{attrs:{label:"1"}},[e._v("当前任务")]),t("el-radio-button",{attrs:{label:"2"}},[e._v("历史任务")])],1),t("el-input",{staticStyle:{width:"300px",margin:"0 5px"},attrs:{placeholder:"关键字查询",clearable:"",size:"mini"},on:{clear:function(t){return e.$refs.table.reload()},blur:function(t){return e.$refs.table.reload()}},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.$refs.table.reload()}},model:{value:e.keyword,callback:function(t){e.keyword=t},expression:"keyword"}},[t("i",{staticClass:"el-input__icon el-icon-search",attrs:{slot:"suffix"},slot:"suffix"})]),t("filterOption",{attrs:{options:e.cond.conditions},on:{change:function(t){return e.$refs.table.reload()}},model:{value:e.filters,callback:function(t){e.filters=t},expression:"filters"}})],1)]},proxy:!0}])})],1)])},_=[],w=i(32268),S=i(95093),C=i.n(S),z=i(83779),$=i(72953),D=i(22562);let A;var T={components:{tableList:w.A,filterOption:$.A},data(){return{columns:[{label:"task_id",prop:"task_id"},{label:"节点ip",prop:"inet"},{label:"broker_name",prop:"broker_name"},{label:"状态",prop:"failed",formatter(e){return e.executed?e.failed?"失败":"成功":"未执行"}},{label:"原因",prop:"reason"},{label:"创建时间",formatter(e){return C()(e.created_at).format("YYYY-MM-DD HH:mm:ss")}}],visible:!1,tableData:[],keyword:null,taskType:"1",timer:null,cond:null,filters:[],pagination:{pageIndex:1,pageSize:15,total:0}}},created(){this.getCond()},methods:{async loadData(){try{const e="1"===this.taskType?"/effect/progresses":"/effect/progress/histories",{data:t}=await u().get(e+"?"+D["default"].parses({current:this.pagination.pageIndex,size:this.pagination.pageSize,keyword:this.keyword?.trim()||"",filters:this.filters}));this.tableData=t?.records||[],this.pagination.total=t.total||0}catch(e){this.$message.error(e)}},domReady(){},async loadProgress(){try{clearTimeout(this.timer);const{data:e}=await u().get("/effect/progress");this.progress=e,this.$nextTick((()=>{this.initEcharts(e)})),this.timer=setTimeout((()=>{this.loadProgress()}),2e3)}catch(e){console.error(e)}},openProgress(){this.loadData(),this.loadProgress(),this.visible=!0},initEcharts({executed:e,failed:t,count:i}){const a={title:{text:"进行中总数："+i,textStyle:{fontSize:12},left:"center",top:"middle"},grid:{top:10},series:[{name:"当前进行任务",type:"pie",radius:["50%","80%"],avoidLabelOverlap:!1,itemStyle:{borderRadius:10,borderColor:"#fff",borderWidth:2},label:{formatter:"{b}: {c}"},emphasis:{label:{show:!0,fontWeight:"bold"}},data:[{value:i-e,name:"未执行"},{value:t,name:"执行失败"},{value:e-t,name:"执行成功"}]}]};A||(A=A||z.init(this.$refs.echartRef),A.on("click",(e=>{}))),A.setOption(a)},close(){clearTimeout(this.timer),A?.dispose(),A=null},async getCond(){try{const{data:e}=await u()("/effect/progress/cond");this.cond=e}catch(e){console.error(e)}}},beforeDestroy(){this.close()}},I=T,R=(0,b.A)(I,x,_,!1,null,"7b3a76fc",null),O=R.exports,P={name:"index",components:{RelationForm:v,Page:k.A,progressDetail:O},data(){return{current:1,size:15,pageTotal:0,relationData:[],isAdd:!0,delLoading:!1,selectInput:"",executed:0,count:0,timer:null,progress:{count:0,executed:0,failed:0},showProgress:!1}},computed:{percent(){return 0===this.progress.count?100:Math.floor(this.progress.executed/this.progress.count*100)}},created(){this.getRelation()},methods:{enableChange(e){var t={id:e.id,enable:e.enable,version:e.version};this.$request.fetchPatchEnable(t).then((()=>{this.getRelation()}))},handleSizeChange(e){this.size=e,this.getRelation()},handleCurrentChange(e){this.current=e,this.getRelation()},addRelation(){const e=this.$refs.forms;this.isAdd=!0,e.dialogVisible=!0,e.form={id:null,name:null,tags:[],compounds:[],substances:[],version:null,enable:!1,exclusion:[]},e.nodeTextarea=""},subDelete(e){this.$refs[e].doClose();var t=this;t.delLoading=!0,this.$request.fetchDelRelation(e).then((e=>{t.$message({message:"删除成功!!!",type:"success"}),t.delLoading=!1,t.getRelation()})).catch((e=>{this.$message.error(e.data)}))},getRelation(){var e={current:this.current,size:this.size,keyword:this.selectInput};this.$request.fetchGetRelation(e).then((e=>{this.relationData=e.data.records,this.pageTotal=e.data.total}))},handleEdit(e,t){this.isAdd=!1;const i=this.$refs.forms;if(i.dialogVisible=!0,null!==t.exclusion)var a=t.exclusion.join("\n");i.form={id:t.id,name:t.name,tags:t.tags,compounds:t.compounds,substances:t.substances?.map((e=>e.id)),version:t.version,enable:t.enable,exclusion:a},i.nodeTextarea=a},progressShow(){},progressHide(){clearTimeout(this.timer)}},beforeDestroy(){clearTimeout(this.timer)}},U=P,B=(0,b.A)(U,a,l,!1,null,null,null),V=B.exports}}]);