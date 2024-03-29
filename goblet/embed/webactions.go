package embed

import (
	"io"
	"strings"
)

const webactionsjs string = `function webactionRequestViaOptions(callbackurl,formid,actiontarget,command){
	postForm({
		url_ref:callbackurl,
		form_ref:formid,
		target:actiontarget,
		command:command
	});
}

var lasturlref="";

function postByElem(elem) {
	var options={}
	$(elem).each(function() {
		$.each(this.attributes, function() {
			if(this.specified) {
				if (this.name=="url_ref" || this.name=="form_ref" || this.name=="progress_elem" || this.name=="target" || this.name=="command"){
					options[this.name] = this.value;
				}
			}
		});
	});
	postForm(options);
}

function postForm(options){
	if(options==undefined) return;
	if(options.url_ref==undefined||options.url_ref==""){
		return;
	}
	
	var hasForm=false;
	
	var progressElem="";
	var errorElem="";
	var urlref="";
	var formid="";
	var command="";
	
	if(options.command!=undefined&&options.command!=""){
		command=options.command;
	}
	
	if(options.form_ref!=undefined&&options.form_ref!=""){
		hasForm=true;
		formid=options.form_ref;
	}
	
	var target="";

	if(options.progress_elem!=undefined){
		progressElem=options.progress_elem+"";
	} else{
		$.blockUI({ 
			message : '<span style="font-size:1.2em" id="showprogress">Please wait ...</span>',
			css: { 
			border: 'none', 
	        padding: '15px', 
	        backgroundColor: '#000', 
	        '-webkit-border-radius': '10px', 
	        '-moz-border-radius': '10px', 
	        opacity: .7, 
	        color: '#fff'
			}
		});
		progressElem="#showprogress";
	}
	
	if(progressElem!=undefined){
		if (progressElem!="#showprogress") {
			$(progressElem).show();
		}
	}
	
	if(options.error_elem!=undefined){
		errorElem=options.error_elem+"";
	}	
	if(options.url_ref!=undefined){
		urlref=options.url_ref+"";
    } else {
    	urlref=lasturlref+"";
    }
    if (hasForm) {
        if(options.form_ref!=undefined){
            formid=options.form_ref+"";
        }
    }
	if(options.target!=undefined){
		target=options.target+"";
	}
    var formData = new FormData();
    var urlparams=getAllUrlParams(urlref);
	if (urlref.indexOf("?")>-1){
		urlref=urlref.slice(0,urlref.indexOf("?"));
		lasturlref=urlref+"";
	} else {
		lasturlref=urlref+"";
	}
	
	if (urlparams!=undefined){
		Object.keys(urlparams).forEach(function(key) {
			if(Array.isArray(urlparams[key])){
				alert("arr:"+key);
				urlparams[key].forEach(function(val){
					formData.append(key,val);
				});
			} else {
				formData.append(key,urlparams[key]);
			}
		});
	}
    var formIds=formid.trim()==""?[]:formid.split("|")
    if (hasForm) {
        formIds.forEach(function(fid,i,arr){
					  if($(fid).length){
							if(!$(fid).is("form")){
								if ($(fid+" select[name],input[name],textarea[name]")!=undefined){
									$(fid+" select[name],input[name],textarea[name]").each(function(){
										var input = $(this); 
										if (input.attr("name")!=""){
											if(input.attr("type")!=undefined && input.attr("type")!="button"&&input.attr("type")!="submit"&&input.attr("type")!="image"){
													if(input.attr("type")=="file"){
															formData.append(input.attr("name"),input[0].files[0]);
													} else {
															formData.append(input.attr("name"),input.val());
													}
											} else {
												formData.append(input.attr("name"),input.val());
											}
										}
									});
								}
								if ($(fid+" textarea")!=undefined) {
									$(fid+" textarea").each(function(){
											var input = $(this);
											if(input.attr("name")!=""){
													formData.append(input.attr("name"),input.text());
											}
									});
								}
							}
            }
        });
    }
	
	if (command!=""){
		formData.append("command",command);
	}
	
    $.ajax({
        xhr: function () {
            var xhr = $.ajaxSettings.xhr();
            xhr.upload.onprogress = function (e) {
            	if(progressElem!=undefined&&progressElem!=""){
            		$(progressElem).html(Math.floor(e.loaded / e.total * 100) + '%');
            	}
            };
            xhr.withCredentials = false;
            return xhr;
        },
        contentType: false,
        processData: false,
        type: 'POST',
        data: formData,
        url: urlref,
        success: function (response,textStatus,xhr) {
			if(xhr.getResponseHeader("Content-Disposition")==null){
				if(progressElem!=undefined){
							if (progressElem=="#showprogress") {
										$.unblockUI();
							} else {
								$(progressElem).hide();
							}
				}
				var parsed=parseActiveString("script||","||script",response);
				var parsedScript=parsed[1].join("");
				response=parsed[0].trim();
				var targets=[];
				var targetSections=[];
				if(response!=""){
					if(response.indexOf("replace-content||")>-1){
						parsed=parseActiveString("replace-content||","||replace-content",response);
						response=parsed[0];
						parsed[1].forEach(function(possibleTargetContent,i){
							if(possibleTargetContent.indexOf("||")>-1){
								targets[targets.length]=[possibleTargetContent.substring(0,possibleTargetContent.indexOf("||")),possibleTargetContent.substring(possibleTargetContent.indexOf("||")+"||".length,possibleTargetContent.length)];
							}        				
						});
					}
					targets.unshift([target,response]);
				}
				if(targets.length>0){
					targets.forEach(function(targetSec){
						if ($(targetSec[0]).length>0) {
									if (targetSec[0].startsWith("#")) {
										$(targetSec[0]).html(targetSec[1]);
									} else {
										$(targetSec[0]).each(function(i){
											$(this).html(targetSec[1])
										});
									}
						}
					});
				}
				if(parsedScript!=""){
					eval(parsedScript);
				}
			} else {
				var contentdisposition=(""+xhr.getResponseHeader("Content-Disposition")).trim();
				if (contentdisposition.indexOf("attachment;")>-1) {
					contentdisposition=contentdisposition.substr(contentdisposition.indexOf("attachment;")+"attachment;".length).trim();
				}
				var contenttype=(""+xhr.getResponseHeader("Content-Type")).trim();
				if (contenttype.indexOf(";")>-1) {
					contenttype=contenttype.substr(0,contenttype.indexOf(";")).trim();
				}
				if (contentdisposition.indexOf("filename=")>-1) {
					contentdisposition=contentdisposition.substr(contentdisposition.indexOf("filename=")+"filename=".length).trim();
					contentdisposition=contentdisposition.replace(/"/i,"")
					contentdisposition=contentdisposition.replace(/"/i,"")
				}
				safeData(responseText,contentdisposition,contenttype);
			}
        },
        error: function(jqXHR, textStatus, textThrow) {
        	if(progressElem!=undefined){
						if (progressElem=="#showprogress") {
									$.unblockUI();
						} else {
							$(progressElem).hide();
						}
        	}
        	if(errorElem!=undefined&&errorElem!=""){
        		$(errorElem).html("Error loading request: "+textStatus);
        	}
        }
    });
}

function safeData(data, fileName,contentType) {
    var a = document.createElement("a");
    document.body.appendChild(a);
    a.style = "display: none";
   
		blob = new Blob([data], {type: contentType}),
		url = window.URL.createObjectURL(blob);
	a.href = url;
	a.download = fileName;
	a.click();
	window.URL.revokeObjectURL(url);    
}

function parseActiveString(labelStart,labelEnd,passiveString){
	this.parsedPassiveString="";
	this.parsedActiveString="";
	this.parsedActiveArr=[];
	this.passiveStringIndex=0;
	this.passiveStringArr=Array.from(passiveString);
	
	this.passiveStringLen=this.passiveStringArr.length;
	
	this.labelStartIndex=0;
	this.labelEndIndex=0;
	
	this.labelStartArr=Array.from(labelStart);
	this.labelEndArr=Array.from(labelEnd);
	this.pc='';
	
	this.passiveStringArr.forEach(function(c,i){
		
		if(this.labelEndIndex==0&&this.labelStartIndex<this.labelStartArr.length){
			if(this.labelStartIndex>0&&this.labelStartArr[this.labelStartIndex-1]==pc&&this.labelStartArr[this.labelStartIndex]!=c){
				this.parsedPassiveString+=labelStart.substring(0,this.labelStartIndex);
				this.labelStartIndex=0;
			}
			if(this.labelStartArr[this.labelStartIndex]==c){
				
				this.labelStartIndex++;
				if(this.labelStartIndex==this.labelStartArr.length){
					
				}
			}
			else{
				if(this.labelStartindex>0){
					this.parsedPassiveString+=labelStart.substring(0,this.labelStartIndex);
					this.labelStartIndex=0;
				}
				this.parsedPassiveString+=(c+"");
			}
		}
		else if(this.labelStartIndex==this.labelStartArr.length&&this.labelEndIndex<this.labelEndArr.length){
			if(this.labelEndArr[this.labelEndIndex]==c){
				this.labelEndIndex++;
				if(this.labelEndIndex==this.labelEndArr.length){
					this.parsedActiveArr[this.parsedActiveArr.length]=this.parsedActiveString+"";
					this.parsedActiveString="";
					this.labelEndIndex=0;
					this.labelStartIndex=0;
				}
			}
			else{
				if(this.labelEndIndex>0){
					this.parsedActiveString+=labelEnd.substring(0,this.labelEndIndex);
					this.labelEndIndex=0;
				}
				this.parsedActiveString+=(c+"");
			}
		}
		
		this.pc=c;		
	});
	return [this.parsedPassiveString,this.parsedActiveArr]
}

function getAllUrlParams(url) {

    // get query string from url (optional) or window
    var queryString = url ? url.split('?')[1] : "";
    
    // we'll store the parameters here
    var obj = {};
  
    // if query string exists
    if (queryString) {
  
      // stuff after # is not part of query string, so get rid of it
      queryString = queryString.split('#')[0];
  
      // split our query string into its component parts
      var arr = queryString.split('&');
  
      for (var i = 0; i < arr.length; i++) {
        // separate the keys and the values
        var a = arr[i].split('=');
  
        // set parameter name and value (use 'true' if empty)
        var paramName = decodeURIComponent(a[0]);
        var paramValue = typeof (a[1]) === 'undefined' ? true : decodeURIComponent(a[1]);
  
        // (optional) keep case consistent
        paramName = paramName.toLowerCase();
        if (typeof paramValue === 'string') paramValue = paramValue.toLowerCase();
  
        // if the paramName ends with square brackets, e.g. colors[] or colors[2]
        if (paramName.match(/\[(\d+)?\]$/)) {
  
          // create key if it doesn't exist
          var key = paramName.replace(/\[(\d+)?\]/, '');
          if (!obj[key]) obj[key] = [];
  
          // if it's an indexed array e.g. colors[2]
          if (paramName.match(/\[\d+\]$/)) {
            // get the index value and add the entry at the appropriate position
            var index = /\[(\d+)\]/.exec(paramName)[1];
            obj[key][index] = paramValue;
          } else {
            // otherwise add the value to the end of the array
            obj[key].push(paramValue);
          }
        } else {
          // we're dealing with a string
          if (!obj[paramName]) {
            // if it doesn't exist, create property
            obj[paramName] = paramValue;
          } else if (obj[paramName] && typeof obj[paramName] === 'string'){
            // if property does exist and it's a string, convert it to an array
            obj[paramName] = [obj[paramName]];
            obj[paramName].push(paramValue);
          } else {
            // otherwise add the property
            obj[paramName].push(paramValue);
          }
        }
      }
    }
    
    return obj;
  }`

const webactionsminjs string = `function webactionRequestViaOptions(r,e,t,s){postForm({url_ref:r,form_ref:e,target:t,command:s})}var lasturlref="";function postByElem(r){var e={};$(r).each(function(){$.each(this.attributes,function(){this.specified&&("url_ref"!=this.name&&"form_ref"!=this.name&&"progress_elem"!=this.name&&"target"!=this.name&&"command"!=this.name||(e[this.name]=this.value))})}),postForm(e)}function postForm(options){if(null!=options&&null!=options.url_ref&&""!=options.url_ref){var hasForm=!1,progressElem="",errorElem="",urlref="",formid="",command="";null!=options.command&&""!=options.command&&(command=options.command),null!=options.form_ref&&""!=options.form_ref&&(hasForm=!0,formid=options.form_ref);var target="";null!=options.progress_elem?progressElem=options.progress_elem+"":($.blockUI({message:'<span style="font-size:1.2em" id="showprogress">Please wait ...</span>',css:{border:"none",padding:"15px",backgroundColor:"#000","-webkit-border-radius":"10px","-moz-border-radius":"10px",opacity:.7,color:"#fff"}}),progressElem="#showprogress"),null!=progressElem&&"#showprogress"!=progressElem&&$(progressElem).show(),null!=options.error_elem&&(errorElem=options.error_elem+""),urlref=null!=options.url_ref?options.url_ref+"":lasturlref+"",hasForm&&null!=options.form_ref&&(formid=options.form_ref+""),null!=options.target&&(target=options.target+"");var formData=new FormData,urlparams=getAllUrlParams(urlref);urlref.indexOf("?")>-1?(urlref=urlref.slice(0,urlref.indexOf("?")),lasturlref=urlref+""):lasturlref=urlref+"",null!=urlparams&&Object.keys(urlparams).forEach(function(r){Array.isArray(urlparams[r])?(alert("arr:"+r),urlparams[r].forEach(function(e){formData.append(r,e)})):formData.append(r,urlparams[r])});var formIds=""==formid.trim()?[]:formid.split("|");hasForm&&formIds.forEach(function(r,e,t){$(r).length&&($(r).is("form")||(null!=$(r+" select[name],input[name],textarea[name]")&&$(r+" select[name],input[name],textarea[name]").each(function(){var r=$(this);""!=r.attr("name")&&(null!=r.attr("type")&&"button"!=r.attr("type")&&"submit"!=r.attr("type")&&"image"!=r.attr("type")&&"file"==r.attr("type")?formData.append(r.attr("name"),r[0].files[0]):formData.append(r.attr("name"),r.val()))}),null!=$(r+" textarea")&&$(r+" textarea").each(function(){var r=$(this);""!=r.attr("name")&&formData.append(r.attr("name"),r.text())})))}),""!=command&&formData.append("command",command),$.ajax({xhr:function(){var r=$.ajaxSettings.xhr();return r.upload.onprogress=function(r){null!=progressElem&&""!=progressElem&&$(progressElem).html(Math.floor(r.loaded/r.total*100)+"%")},r.withCredentials=!1,r},contentType:!1,processData:!1,type:"POST",data:formData,url:urlref,success:function(response){null!=progressElem&&("#showprogress"==progressElem?$.unblockUI():$(progressElem).hide());var parsed=parseActiveString("script||","||script",response),parsedScript=parsed[1].join("");response=parsed[0].trim();var targets=[],targetSections=[];""!=response&&(response.indexOf("replace-content||")>-1&&(parsed=parseActiveString("replace-content||","||replace-content",response),response=parsed[0],parsed[1].forEach(function(r,e){r.indexOf("||")>-1&&(targets[targets.length]=[r.substring(0,r.indexOf("||")),r.substring(r.indexOf("||")+"||".length,r.length)])})),targets.unshift([target,response])),targets.length>0&&targets.forEach(function(r){$(r[0]).length>0&&(r[0].startsWith("#")?$(r[0]).html(r[1]):$(r[0]).each(function(e){$(this).html(r[1])}))}),""!=parsedScript&&eval(parsedScript)},error:function(r,e,t){alert(t),null!=progressElem&&("#showprogress"==progressElem?$.unblockUI():$(progressElem).hide()),null!=errorElem&&""!=errorElem&&$(errorElem).html("Error loading request: "+e)}})}}function parseActiveString(r,e,t){return this.parsedPassiveString="",this.parsedActiveString="",this.parsedActiveArr=[],this.passiveStringIndex=0,this.passiveStringArr=Array.from(t),this.passiveStringLen=this.passiveStringArr.length,this.labelStartIndex=0,this.labelEndIndex=0,this.labelStartArr=Array.from(r),this.labelEndArr=Array.from(e),this.pc="",this.passiveStringArr.forEach(function(t,s){0==this.labelEndIndex&&this.labelStartIndex<this.labelStartArr.length?(this.labelStartIndex>0&&this.labelStartArr[this.labelStartIndex-1]==pc&&this.labelStartArr[this.labelStartIndex]!=t&&(this.parsedPassiveString+=r.substring(0,this.labelStartIndex),this.labelStartIndex=0),this.labelStartArr[this.labelStartIndex]==t?(this.labelStartIndex++,this.labelStartIndex,this.labelStartArr.length):(this.labelStartindex>0&&(this.parsedPassiveString+=r.substring(0,this.labelStartIndex),this.labelStartIndex=0),this.parsedPassiveString+=t+"")):this.labelStartIndex==this.labelStartArr.length&&this.labelEndIndex<this.labelEndArr.length&&(this.labelEndArr[this.labelEndIndex]==t?(this.labelEndIndex++,this.labelEndIndex==this.labelEndArr.length&&(this.parsedActiveArr[this.parsedActiveArr.length]=this.parsedActiveString+"",this.parsedActiveString="",this.labelEndIndex=0,this.labelStartIndex=0)):(this.labelEndIndex>0&&(this.parsedActiveString+=e.substring(0,this.labelEndIndex),this.labelEndIndex=0),this.parsedActiveString+=t+"")),this.pc=t}),[this.parsedPassiveString,this.parsedActiveArr]}function getAllUrlParams(r){var e=r?r.split("?")[1]:"",t={};if(e)for(var s=(e=e.split("#")[0]).split("&"),a=0;a<s.length;a++){var n=s[a].split("="),l=decodeURIComponent(n[0]),i=void 0===n[1]||decodeURIComponent(n[1]);if(l=l.toLowerCase(),"string"==typeof i&&(i=i.toLowerCase()),l.match(/\[(\d+)?\]$/)){var o=l.replace(/\[(\d+)?\]/,"");if(t[o]||(t[o]=[]),l.match(/\[\d+\]$/)){var p=/\[(\d+)\]/.exec(l)[1];t[o][p]=i}else t[o].push(i)}else t[l]?t[l]&&"string"==typeof t[l]?(t[l]=[t[l]],t[l].push(i)):t[l].push(i):t[l]=i}return t}`

//WebactionsJS WebactionsJS
func WebactionsJS(min ...bool) io.Reader {
	if len(min) == 1 && min[0] {
		return strings.NewReader(webactionsminjs)
	}
	return strings.NewReader(webactionsjs)
}
