package embed

import (
	"io"
	"strings"
)

const hcoffcanvasnavcss string = `html.hc-nav-yscroll{overflow-y:scroll}body.hc-nav-open{overflow:visible;position:fixed;width:100%;min-height:100%}.hc-offcanvas-nav{visibility:hidden;display:none;position:fixed;top:0;height:100%;z-index:9999}.hc-offcanvas-nav.is-ios *{cursor:pointer !important}.hc-offcanvas-nav .nav-container{position:fixed;z-index:9998;top:0;width:240px;height:100%;max-width:100%;max-height:100%;box-sizing:border-box;transition:-webkit-transform .4s ease;transition:transform .4s ease;transition:transform .4s ease, -webkit-transform .4s ease}.hc-offcanvas-nav .nav-wrapper{width:100%;height:100%;max-height:100vh;-ms-scroll-chaining:none;overscroll-behavior:none;box-sizing:border-box}.hc-offcanvas-nav .nav-content{height:100%;max-height:100vh}.hc-offcanvas-nav .nav-wrapper-0>.nav-content{overflow:scroll;overflow-x:visible;overflow-y:auto;box-sizing:border-box}.hc-offcanvas-nav ul{list-style:none;margin:0;padding:0}.hc-offcanvas-nav li{position:relative;display:block}.hc-offcanvas-nav li.level-open>.nav-wrapper{visibility:visible}.hc-offcanvas-nav input[type="checkbox"]{display:none}.hc-offcanvas-nav label{position:absolute;top:0;left:0;right:0;bottom:0;z-index:10;cursor:pointer}.hc-offcanvas-nav a{position:relative;display:block;box-sizing:border-box;cursor:pointer}.hc-offcanvas-nav a,.hc-offcanvas-nav a:hover{text-decoration:none}.hc-offcanvas-nav .nav-item{position:relative;display:block;box-sizing:border-box}.hc-offcanvas-nav.disable-body::after,.hc-offcanvas-nav .nav-wrapper::after{content:'';position:fixed;z-index:9990;top:0;left:0;right:0;bottom:0;width:100%;height:100%;-ms-scroll-chaining:none;overscroll-behavior:none;visibility:hidden;opacity:0;transition:visibility 0s ease .4s,opacity .4s ease}.hc-offcanvas-nav.disable-body.nav-open::after,.hc-offcanvas-nav .sub-level-open::after{visibility:visible;opacity:1;transition-delay:.05s}.hc-offcanvas-nav:not(.nav-open)::after{pointer-events:none}.hc-offcanvas-nav.nav-levels-expand .nav-content{overflow:scroll;overflow-x:visible;overflow-y:auto;box-sizing:border-box}.hc-offcanvas-nav.nav-levels-expand .nav-wrapper::after{display:none}.hc-offcanvas-nav.nav-levels-expand ul .nav-wrapper{min-width:0;max-height:0;overflow:hidden;transition:height 0s ease .4s}.hc-offcanvas-nav.nav-levels-expand .level-open>.nav-wrapper{max-height:none}.hc-offcanvas-nav.nav-levels-overlap .nav-content{overflow:scroll;overflow-x:visible;overflow-y:auto;box-sizing:border-box}.hc-offcanvas-nav.nav-levels-overlap ul .nav-wrapper{position:absolute;z-index:9999;top:0;height:100%;visibility:hidden;transition:visibility 0s ease .4s,-webkit-transform .4s ease;transition:visibility 0s ease .4s,transform .4s ease;transition:visibility 0s ease .4s,transform .4s ease,-webkit-transform .4s ease}.hc-offcanvas-nav.nav-levels-overlap ul li.nav-parent{position:static}.hc-offcanvas-nav.nav-levels-overlap ul li.level-open>.nav-wrapper{visibility:visible;-webkit-transform:translate3d(0, 0, 0);transform:translate3d(0, 0, 0);transition:-webkit-transform .4s ease;transition:transform .4s ease;transition:transform .4s ease, -webkit-transform .4s ease}.hc-offcanvas-nav.nav-position-left{left:0}.hc-offcanvas-nav.nav-position-left .nav-container{left:0;-webkit-transform:translate3d(-240px, 0, 0);transform:translate3d(-240px, 0, 0)}.hc-offcanvas-nav.nav-position-left.nav-levels-overlap li .nav-wrapper{left:0;-webkit-transform:translate3d(-100%, 0, 0);transform:translate3d(-100%, 0, 0)}.hc-offcanvas-nav.nav-position-right{right:0}.hc-offcanvas-nav.nav-position-right .nav-container{right:0;-webkit-transform:translate3d(240px, 0, 0);transform:translate3d(240px, 0, 0)}.hc-offcanvas-nav.nav-position-right.nav-levels-overlap li .nav-wrapper{right:0;-webkit-transform:translate3d(100%, 0, 0);transform:translate3d(100%, 0, 0)}.hc-offcanvas-nav.nav-position-top{top:0}.hc-offcanvas-nav.nav-position-top .nav-container{top:0;width:100%;height:auto;-webkit-transform:translate3d(0, -100%, 0);transform:translate3d(0, -100%, 0)}.hc-offcanvas-nav.nav-position-top.nav-levels-overlap li .nav-wrapper{left:0;-webkit-transform:translate3d(0, -100%, 0);transform:translate3d(0, -100%, 0)}.hc-offcanvas-nav.nav-position-bottom{top:auto;bottom:0}.hc-offcanvas-nav.nav-position-bottom .nav-container{top:auto;bottom:0;width:100%;height:auto;-webkit-transform:translate3d(0, 100%, 0);transform:translate3d(0, 100%, 0)}.hc-offcanvas-nav.nav-position-bottom.nav-levels-overlap li .nav-wrapper{left:0;-webkit-transform:translate3d(0, 100%, 0);transform:translate3d(0, 100%, 0)}.hc-offcanvas-nav.nav-open[class*='hc-nav-'] div.nav-container{-webkit-transform:translate3d(0, 0, 0);transform:translate3d(0, 0, 0)}.hc-nav-trigger{position:absolute;cursor:pointer;-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none;display:none;top:20px;z-index:9980;width:30px;min-height:24px}.hc-nav-trigger span{width:30px;top:50%;-webkit-transform:translateY(-50%);transform:translateY(-50%);-webkit-transform-origin:50% 50%;transform-origin:50% 50%}.hc-nav-trigger span,.hc-nav-trigger span::before,.hc-nav-trigger span::after{display:block;position:absolute;left:0;height:4px;background:#34495E;transition:all .2s ease}.hc-nav-trigger span::before,.hc-nav-trigger span::after{content:'';width:100%}.hc-nav-trigger span::before{top:-10px}.hc-nav-trigger span::after{bottom:-10px}.hc-nav-trigger.toggle-open span{background:rgba(0,0,0,0);-webkit-transform:rotate(45deg);transform:rotate(45deg)}.hc-nav-trigger.toggle-open span::before{-webkit-transform:translate3d(0, 10px, 0);transform:translate3d(0, 10px, 0)}.hc-nav-trigger.toggle-open span::after{-webkit-transform:rotate(-90deg) translate3d(10px, 0, 0);transform:rotate(-90deg) translate3d(10px, 0, 0)}.hc-offcanvas-nav::after,.hc-offcanvas-nav .nav-wrapper::after{background:rgba(0,0,0,0.3)}.hc-offcanvas-nav .nav-container,.hc-offcanvas-nav .nav-wrapper,.hc-offcanvas-nav ul{background:#336ca6}.hc-offcanvas-nav h2{font-size:19px;font-weight:normal;text-align:left;padding:20px 17px;color:#1b3958}.hc-offcanvas-nav a,.hc-offcanvas-nav .nav-item{padding:14px 17px;font-size:15px;color:#fff;z-index:1;background:rgba(0,0,0,0);border-bottom:1px solid #2c5d8f}.hc-offcanvas-nav:not(.touch-device) a:hover{background:#31679e}.hc-offcanvas-nav ul:first-of-type:not(:first-child)>li:first-child:not(.nav-back):not(.nav-close)>a{border-top:1px solid #2c5d8f;margin-top:-1px}.hc-offcanvas-nav li{text-align:left}.hc-offcanvas-nav li.nav-close a,.hc-offcanvas-nav li.nav-back a{background:#2c5d8f;border-top:1px solid #295887;border-bottom:1px solid #295887}.hc-offcanvas-nav li.nav-close a:hover,.hc-offcanvas-nav li.nav-back a:hover{background:#2b5c8d}.hc-offcanvas-nav li.nav-close:not(:first-child) a,.hc-offcanvas-nav li.nav-back:not(:first-child) a{margin-top:-1px}.hc-offcanvas-nav li.nav-parent .nav-item{padding-right:58px}.hc-offcanvas-nav li.nav-close span,.hc-offcanvas-nav li.nav-parent span.nav-next,.hc-offcanvas-nav li.nav-back span{width:45px;position:absolute;top:0;right:0;bottom:0;text-align:center;cursor:pointer;transition:background .2s ease}.hc-offcanvas-nav li.nav-close span::before,.hc-offcanvas-nav li.nav-close span::after{content:'';position:absolute;top:50%;left:50%;width:6px;height:6px;margin-top:-3px;border-top:2px solid #fff;border-left:2px solid #fff}.hc-offcanvas-nav li.nav-close span::before{margin-left:-9px;-webkit-transform:rotate(135deg);transform:rotate(135deg)}.hc-offcanvas-nav li.nav-close span::after{-webkit-transform:rotate(-45deg);transform:rotate(-45deg)}.hc-offcanvas-nav a[href]:not([href="#"])>span.nav-next{border-left:1px solid #2c5d8f}.hc-offcanvas-nav span.nav-next::before,.hc-offcanvas-nav li.nav-back span::before{content:'';position:absolute;top:50%;left:50%;width:8px;height:8px;margin-left:-2px;box-sizing:border-box;border-top:2px solid #fff;border-left:2px solid #fff;-webkit-transform-origin:center;transform-origin:center}.hc-offcanvas-nav span.nav-next::before{-webkit-transform:translate(-50%, -50%) rotate(135deg);transform:translate(-50%, -50%) rotate(135deg)}.hc-offcanvas-nav li.nav-back span::before{-webkit-transform:translate(-50%, -50%) rotate(-45deg);transform:translate(-50%, -50%) rotate(-45deg)}.hc-offcanvas-nav.nav-position-left.nav-open .nav-wrapper{box-shadow:1px 0 2px rgba(0,0,0,0.2)}.hc-offcanvas-nav.nav-position-right.nav-open .nav-wrapper{box-shadow:-1px 0 2px rgba(0,0,0,0.2)}.hc-offcanvas-nav.nav-position-right span.nav-next::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(-45deg);transform:translate(-50%, -50%) rotate(-45deg)}.hc-offcanvas-nav.nav-position-right li.nav-back span::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(135deg);transform:translate(-50%, -50%) rotate(135deg)}.hc-offcanvas-nav.nav-position-top.nav-open .nav-wrapper{box-shadow:0 1px 2px rgba(0,0,0,0.2)}.hc-offcanvas-nav.nav-position-top span.nav-next::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(-135deg);transform:translate(-50%, -50%) rotate(-135deg)}.hc-offcanvas-nav.nav-position-top li.nav-back span::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(45deg);transform:translate(-50%, -50%) rotate(45deg)}.hc-offcanvas-nav.nav-position-bottom.nav-open .nav-wrapper{box-shadow:0 -1px 2px rgba(0,0,0,0.2)}.hc-offcanvas-nav.nav-position-bottom span.nav-next::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(45deg);transform:translate(-50%, -50%) rotate(45deg)}.hc-offcanvas-nav.nav-position-bottom li.nav-back span::before{margin-left:0;margin-right:-2px;-webkit-transform:translate(-50%, -50%) rotate(-135deg);transform:translate(-50%, -50%) rotate(-135deg)}.hc-offcanvas-nav.nav-levels-expand .nav-container ul .nav-wrapper,.hc-offcanvas-nav.nav-levels-none .nav-container ul .nav-wrapper{box-shadow:none;background:transparent}.hc-offcanvas-nav.nav-levels-expand .nav-container ul h2,.hc-offcanvas-nav.nav-levels-none .nav-container ul h2{display:none}.hc-offcanvas-nav.nav-levels-expand .nav-container ul ul .nav-item,.hc-offcanvas-nav.nav-levels-none .nav-container ul ul .nav-item{font-size:14px}.hc-offcanvas-nav.nav-levels-expand .nav-container li,.hc-offcanvas-nav.nav-levels-none .nav-container li{transition:background .3s ease}.hc-offcanvas-nav.nav-levels-expand .nav-container li.level-open,.hc-offcanvas-nav.nav-levels-none .nav-container li.level-open{background:#2e6296}.hc-offcanvas-nav.nav-levels-expand .nav-container li.level-open a,.hc-offcanvas-nav.nav-levels-none .nav-container li.level-open a{border-bottom:1px solid #295887}.hc-offcanvas-nav.nav-levels-expand .nav-container li.level-open a:hover,.hc-offcanvas-nav.nav-levels-none .nav-container li.level-open a:hover{background:#2f649a}.hc-offcanvas-nav.nav-levels-expand .nav-container li.level-open>.nav-item .nav-next::before,.hc-offcanvas-nav.nav-levels-none .nav-container li.level-open>.nav-item .nav-next::before{margin-top:2px;-webkit-transform:translate(-50%, -50%) rotate(45deg);transform:translate(-50%, -50%) rotate(45deg)}.hc-offcanvas-nav.nav-levels-expand .nav-container span.nav-next::before,.hc-offcanvas-nav.nav-levels-none .nav-container span.nav-next::before{margin-top:-2px;-webkit-transform:translate(-50%, -50%) rotate(-135deg);transform:translate(-50%, -50%) rotate(-135deg)}`

const hcoffcanvasnavjs string = `/*
* HC Off-canvas Nav
* ===================
* Version: 3.4.1
* Author: Some Web Media
* Author URL: http://somewebmedia.com
* Plugin URL: https://github.com/somewebmedia/hc-offcanvas-nav
* Description: jQuery plugin for creating off-canvas multi-level navigations
* License: MIT
*/
"use strict";function _typeof(n){return(_typeof="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(n){return typeof n}:function(n){return n&&"function"==typeof Symbol&&n.constructor===Symbol&&n!==Symbol.prototype?"symbol":typeof n})(n)}!function(_,n){var c,L=n.document,Q=_(L.getElementsByTagName("html")[0]),U=(_(L),(/iPad|iPhone|iPod/.test(navigator.userAgent)||!!navigator.platform&&/iPad|iPhone|iPod/.test(navigator.platform))&&!n.MSStream),z="ontouchstart"in n||navigator.maxTouchPoints||n.DocumentTouch&&L instanceof DocumentTouch,V=function(n){return!isNaN(parseFloat(n))&&isFinite(n)},G=function(n){return n.stopPropagation()},I=function(e){return function(n){n.preventDefault(),n.stopPropagation(),"function"==typeof e&&e()}},J=function(n,e,t){var a=t.children(),o=a.length,c=-1<e?Math.max(0,Math.min(e-1,o)):Math.max(0,Math.min(o+e+1,o));0===c?t.prepend(n):a.eq(c-1).after(n)},K=function(n){return-1!==["left","right"].indexOf(n)?"x":"y"},R=(c=function(n){var e=["Webkit","Moz","Ms","O"],t=(L.body||L.documentElement).style,a=n.charAt(0).toUpperCase()+n.slice(1);if(void 0!==t[n])return n;for(var o=0;o<e.length;o++)if(void 0!==t[e[o]+a])return e[o]+a;return!1}("transform"),function(n,e,t){if(c)if(0===e)n.css(c,"");else if("x"===K(t)){var a="left"===t?e:-e;n.css(c,a?"translate3d(".concat(a,"px,0,0)"):"")}else{var o="top"===t?e:-e;n.css(c,o?"translate3d(0,".concat(o,"px,0)"):"")}else n.css(t,e)}),e=function(n,e,t){console.warn("%cHC Off-canvas Nav:%c "+t+"%c '"+n+"'%c is now deprecated and will be removed. Use%c '"+e+"'%c instead.","color: #fa253b","color: default","color: #5595c6","color: default","color: #5595c6","color: default")},X=0;_.fn.extend({hcOffcanvasNav:function(){var n=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};if(!this.length)return this;var A=this,D=_(L.body);n.side&&(e("side","position","option"),n.position=n.side);var H=_.extend({},{maxWidth:1024,pushContent:!1,position:"left",levelOpen:"overlap",levelSpacing:40,levelTitles:!1,navTitle:null,navClass:"",disableBody:!0,closeOnClick:!0,customToggle:null,insertClose:!0,insertBack:!0,labelClose:"Close",labelBack:"Back"},n),E=[],$="nav-open",F=function(n){if(!E.length)return!1;var e=!1;"string"==typeof n&&(n=[n]);for(var t=n.length,a=0;a<t;a++)-1!==E.indexOf(n[a])&&(e=!0);return e};return this.each(function(){var e=_(this);if(e.find("ul").addBack("ul").length){var n,o,c,i,a,s,b="hc-nav-".concat(++X),l=(n="hc-offcanvas-".concat(X,"-style"),o=_('<style id="'.concat(n,'">')).appendTo(_("head")),c={},i={},a=function(n){return";"!==n.substr(-1)&&(n+=";"!==n.substr(-1)?";":""),n},{reset:function(){c={},i={}},add:function(n,e,t){n=n.trim(),e=e.trim(),t?(t=t.trim(),i[t]=i[t]||{},i[t][n]=a(e)):c[n]=a(e)},remove:function(n,e){n=n.trim(),e?(e=e.trim(),void 0!==i[e][n]&&delete i[e][n]):void 0!==c[n]&&delete c[n]},insert:function(){var n="";for(var e in i){for(var t in n+="@media screen and (".concat(e,") {\n"),i[e])n+="".concat(t," { ").concat(i[e][t]," }\n");n+="}\n"}for(var a in c)n+="".concat(a," { ").concat(c[a]," }\n");o.html(n)}});e.addClass("hc-nav ".concat(b));var t,r,p,d=_("<nav>").on("click",G),v=_('<div class="nav-container">').appendTo(d),f=null,u={},h=!1,m=0,g=0,y=0,C=null,x={},k=[];H.customToggle?s=_(H.customToggle).addClass("hc-nav-trigger ".concat(b)).on("click",q):(s=_('<a class="hc-nav-trigger '.concat(b,'"><span></span></a>')).on("click",q),e.after(s));var O=function(){v.css("transition","none"),g=v.outerWidth(),y=v.outerHeight(),l.add(".hc-offcanvas-nav.".concat(b,".nav-position-left .nav-container"),"transform: translate3d(-".concat(g,"px, 0, 0)")),l.add(".hc-offcanvas-nav.".concat(b,".nav-position-right .nav-container"),"transform: translate3d(".concat(g,"px, 0, 0)")),l.add(".hc-offcanvas-nav.".concat(b,".nav-position-top .nav-container"),"transform: translate3d(0, -".concat(y,"px, 0)")),l.add(".hc-offcanvas-nav.".concat(b,".nav-position-bottom .nav-container"),"transform: translate3d(0, ".concat(y,"px, 0)")),l.insert(),v.css("transition",""),T()},T=function(){var n;t=v.css("transition-property").split(",")[0],n=v.css("transition-duration").split(",")[0],r=parseFloat(n)*(/\ds$/.test(n)?1e3:1),p=v.css("transition-timing-function").split(",")[0],H.pushContent&&f&&t&&l.add(function n(e){return"string"==typeof e?e:e.attr("id")?"#"+e.attr("id"):e.attr("class")?e.prop("tagName").toLowerCase()+"."+e.attr("class").replace(/\s+/g,"."):n(e.parent())+" "+e.prop("tagName").toLowerCase()}(H.pushContent),"transition: ".concat(t," ").concat(r,"ms ").concat(p)),l.insert()},S=function(n){var e=s.css("display"),t=!!H.maxWidth&&"max-width: ".concat(H.maxWidth-1,"px");F("maxWidth")&&l.reset(),l.add(".hc-offcanvas-nav.".concat(b),"display: block",t),l.add(".hc-nav-trigger.".concat(b),"display: ".concat(e&&"none"!==e?e:"block"),t),l.add(".hc-nav.".concat(b),"display: none",t),l.add(".hc-offcanvas-nav.".concat(b,".nav-levels-overlap.nav-position-left li.level-open > .nav-wrapper"),"transform: translate3d(-".concat(H.levelSpacing,"px,0,0)"),t),l.add(".hc-offcanvas-nav.".concat(b,".nav-levels-overlap.nav-position-right li.level-open > .nav-wrapper"),"transform: translate3d(".concat(H.levelSpacing,"px,0,0)"),t),l.add(".hc-offcanvas-nav.".concat(b,".nav-levels-overlap.nav-position-top li.level-open > .nav-wrapper"),"transform: translate3d(0,-".concat(H.levelSpacing,"px,0)"),t),l.add(".hc-offcanvas-nav.".concat(b,".nav-levels-overlap.nav-position-bottom li.level-open > .nav-wrapper"),"transform: translate3d(0,".concat(H.levelSpacing,"px,0)"),t),l.insert(),(!n||n&&F("pushContent"))&&("string"==typeof H.pushContent?(f=_(H.pushContent)).length||(f=null):f=H.pushContent instanceof jQuery?H.pushContent:null),v.css("transition","none");var a=d.hasClass($),o=["hc-offcanvas-nav",H.navClass||"",b,H.navClass||"","nav-levels-"+H.levelOpen||"none","nav-position-"+H.position,H.disableBody?"disable-body":"",U?"is-ios":"",z?"touch-device":"",a?$:""].join(" ");d.off("click").attr("class","").addClass(o),H.disableBody&&d.on("click",j),n?O():setTimeout(O,1)},w=function(){var n;u=function c(n){var e=[];return n.each(function(){var n=_(this),o={classes:n.attr("class"),items:[]};n.children("li").each(function(){var n=_(this),e=n.children().filter(function(){var n=_(this);return n.is(":not(ul)")&&!n.find("ul").length}).add(n.contents().filter(function(){return 3===this.nodeType&&this.nodeValue.trim()})),t=n.find("ul"),a=t.first().add(t.first().siblings("ul"));a.length&&!n.data("hc-uniqid")&&n.data("hc-uniqid",Math.random().toString(36).substr(2)+"-"+Math.random().toString(36).substr(2)),o.items.push({uniqid:n.data("hc-uniqid"),classes:n.attr("class"),$content:e,subnav:a.length?c(a):[]})}),e.push(o)}),e}((n=e.find("ul").addBack("ul")).first().add(n.first().siblings("ul")))},B=function(n){n&&(v.empty(),x={}),function h(n,e,m,t,a){var g=_('<div class="nav-wrapper nav-wrapper-'.concat(m,'">')).appendTo(e).on("click",G),o=_('<div class="nav-content">').appendTo(g);if(t&&o.prepend("<h2>".concat(t,"</h2>")),_.each(n,function(n,e){var u=_("<ul>").addClass(e.classes).appendTo(o);_.each(e.items,function(n,e){var t=e.$content,a=t.find("a").addBack("a"),o=a.length?a.clone(!0,!0).addClass("nav-item"):_('<span class="nav-item">').append(t.clone(!0,!0)).on("click",G);a.length&&o.on("click",function(n){n.stopPropagation(),a[0].click()}),"#"===o.attr("href")&&o.on("click",function(n){n.preventDefault()}),H.closeOnClick&&(!1===H.levelOpen||"none"===H.levelOpen?o.filter("a").filter('[data-nav-close!="false"]').on("click",j):o.filter("a").filter('[data-nav-close!="false"]').filter(function(){var n=_(this);return!e.subnav.length||n.attr("href")&&"#"!==n.attr("href").charAt(0)}).on("click",j));var c=_("<li>").addClass(e.classes).append(o);if(u.append(c),H.levelSpacing&&("expand"===H.levelOpen||!1===H.levelOpen||"none"===H.levelOpen)){var i=H.levelSpacing*m;i&&u.css("text-indent","".concat(i,"px"))}if(e.subnav.length){var s=m+1,l=e.uniqid,r="";if(x[s]||(x[s]=0),c.addClass("nav-parent"),!1!==H.levelOpen&&"none"!==H.levelOpen){var p=x[s],d=_('<span class="nav-next">').appendTo(o),v=_('<label for="'.concat(b,"-").concat(s,"-").concat(p,'">')).on("click",G),f=_('<input type="checkbox" id="'.concat(b,"-").concat(s,"-").concat(p,'">')).attr("data-level",s).attr("data-index",p).val(l).on("click",G).on("change",M);-1!==k.indexOf(l)&&(g.addClass("sub-level-open").on("click",function(){return W(s,p)}),c.addClass("level-open"),f.prop("checked",!0)),c.prepend(f),r=!0===H.levelTitles?t.text().trim():"",o.attr("href")&&"#"!==o.attr("href").charAt(0)?d.append(v):o.prepend(v.on("click",function(){_(this).parent().trigger("click")}))}x[s]++,h(e.subnav,c,s,r,x[s]-1)}})}),m&&void 0!==a&&!1!==H.insertBack&&"overlap"===H.levelOpen){var c=o.children("ul"),i=_('<li class="nav-back"><a href="#">'.concat(H.labelBack||"","<span></span></a></li>"));i.children("a").on("click",I(function(){return W(m,a)})),!0===H.insertBack?c.first().prepend(i):V(H.insertBack)&&J(i,H.insertBack,c)}if(0===m&&!1!==H.insertClose){var s=o.children("ul"),l=_('<li class="nav-close"><a href="#">'.concat(H.labelClose||"","<span></span></a></li>"));l.children("a").on("click",I(j)),!0===H.insertClose?s.first().prepend(l):V(H.insertClose)&&J(l,H.insertClose,s.first().add(s.first().siblings("ul")))}}(u,v,0,H.navTitle)};S(),w(),B(),D.append(d);var P=function(n,e,t){var a=_("#".concat(b,"-").concat(n,"-").concat(e)),o=a.val(),c=a.parent("li"),i=c.closest(".nav-wrapper");if(a.prop("checked",!1),i.removeClass("sub-level-open"),c.removeClass("level-open"),-1!==k.indexOf(o)&&k.splice(k.indexOf(o),1),t&&"overlap"===H.levelOpen&&(i.off("click").on("click",G),R(v,(n-1)*H.levelSpacing,H.position),f)){var s="x"===K(H.position)?g:y;R(f,s+(n-1)*H.levelSpacing,H.position)}};A.settings=function(n){return n?H[n]:Object.assign({},H)},A.isOpen=function(){return d.hasClass($)},A.open=N,A.close=j,A.update=function(n,e){if(E=[],"object"===_typeof(n)){for(var t in n)H[t]!==n[t]&&E.push(t);H=_.extend({},H,n),S(!0),B(!0)}(!0===n||e)&&(S(!0),w(),B(!0))}}else console.error("%c! HC Offcanvas Nav:%c Menu must contain <ul> element.","color: #fa253b","color: default");function M(){var n=_(this),e=Number(n.attr("data-level")),t=Number(n.attr("data-index"));n.prop("checked")?function(n,e){var t=_("#".concat(b,"-").concat(n,"-").concat(e)),a=t.val(),o=t.parent("li"),c=o.closest(".nav-wrapper");if(c.addClass("sub-level-open"),o.addClass("level-open"),-1===k.indexOf(a)&&k.push(a),"overlap"===H.levelOpen&&(c.on("click",function(){return W(n,e)}),R(v,n*H.levelSpacing,H.position),f)){var i="x"===K(H.position)?g:y;R(f,i+n*H.levelSpacing,H.position)}}(e,t):W(e,t)}function N(){if(h=!0,d.css("visibility","visible").addClass($),s.addClass("toggle-open"),"expand"===H.levelOpen&&C&&clearTimeout(C),H.disableBody&&(m=Q.scrollTop()||D.scrollTop(),L.documentElement.scrollHeight>L.documentElement.clientHeight&&Q.addClass("hc-nav-yscroll"),D.addClass("hc-nav-open"),m&&D.css("top",-m)),f){var n="x"===K(H.position)?g:y;R(f,n,H.position)}setTimeout(function(){A.trigger("open",_.extend({},H))},r+1)}function j(){h=!1,f&&R(f,0),d.removeClass($),v.removeAttr("style"),s.removeClass("toggle-open"),"expand"===H.levelOpen&&-1!==["top","bottom"].indexOf(H.position)?W(0):!1!==H.levelOpen&&"none"!==H.levelOpen&&(C=setTimeout(function(){W(0)},"expand"===H.levelOpen?r:0)),H.disableBody&&(D.removeClass("hc-nav-open"),Q.removeClass("hc-nav-yscroll"),m&&(D.css("top","").scrollTop(m),Q.scrollTop(m),m=0)),setTimeout(function(){d.css("visibility",""),A.trigger("close.$",_.extend({},H)),A.trigger("close.once",_.extend({},H)).off("close.once")},r+1)}function q(n){n.preventDefault(),n.stopPropagation(),h?j():N()}function W(n,e){for(var t=n;t<=Object.keys(x).length;t++)if(t==n&&void 0!==e)P(n,e,!0);else for(var a=0;a<x[t];a++)P(t,a,t==n)}})}})}(jQuery,"undefined"!=typeof window?window:this);`

const hcstickyjs string = `/*!
* HC-Sticky
* =========
* Version: 2.2.3
* Author: Some Web Media
* Author URL: http://somewebmedia.com
* Plugin URL: https://github.com/somewebmedia/hc-sticky
* Description: Cross-browser plugin that makes any element on your page visible while you scroll
* License: MIT
*/
!function(t,e){"use strict";if("object"===("undefined"==typeof module?"undefined":_typeof(module))&&"object"===_typeof(module.exports)){if(!t.document)throw new Error("HC-Sticky requires a browser to run.");module.exports=e(t)}else"function"==typeof define&&define.amd?define("hcSticky",[],e(t)):e(t)}("undefined"!=typeof window?window:this,function(_){"use strict";var U={top:0,bottom:0,bottomEnd:0,innerTop:0,innerSticker:null,stickyClass:"sticky",stickTo:null,followScroll:!0,responsive:null,mobileFirst:!1,onStart:null,onStop:null,onBeforeResize:null,onResize:null,resizeDebounce:100,disable:!1,queries:null,queryFlow:"down"},Y=function(t,e,o){console.warn("%cHC Sticky:%c "+o+"%c '"+t+"'%c is now deprecated and will be removed. Use%c '"+e+"'%c instead.","color: #fa253b","color: default","color: #5595c6","color: default","color: #5595c6","color: default")},$=_.document,Q=function(i){var o=this,f=1<arguments.length&&void 0!==arguments[1]?arguments[1]:{};if("string"==typeof i&&(i=$.querySelector(i)),!i)return!1;f.queries&&Y("queries","responsive","option"),f.queryFlow&&Y("queryFlow","mobileFirst","option");var p={},u=Q.Helpers,s=i.parentNode;"static"===u.getStyle(s,"position")&&(s.style.position="relative");var r,l,a,c,d,y,m,g,h,b,v,S,w,k,E,x,L,T,j,O=function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};u.isEmptyObject(t)&&!u.isEmptyObject(p)||(p=Object.assign({},U,p,t))},t=function(){return p.disable},e=function(){var t,e=p.responsive||p.queries;if(e){var o=_.innerWidth;if(t=f,(p=Object.assign({},U,t||{})).mobileFirst)for(var n in e)n<=o&&!u.isEmptyObject(e[n])&&O(e[n]);else{var i=[];for(var s in e){var r={};r[s]=e[s],i.push(r)}for(var l=i.length-1;0<=l;l--){var a=i[l],c=Object.keys(a)[0];o<=c&&!u.isEmptyObject(a[c])&&O(a[c])}}}},C={css:{},position:null,stick:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};u.hasClass(i,p.stickyClass)||(!1===z.isAttached&&z.attach(),C.position="fixed",i.style.position="fixed",i.style.left=z.offsetLeft+"px",i.style.width=z.width,void 0===t.bottom?i.style.bottom="auto":i.style.bottom=t.bottom+"px",void 0===t.top?i.style.top="auto":i.style.top=t.top+"px",i.classList?i.classList.add(p.stickyClass):i.className+=" "+p.stickyClass,p.onStart&&p.onStart.call(i,Object.assign({},p)))},release:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};if(t.stop=t.stop||!1,!0===t.stop||"fixed"===C.position||null===C.position||!(void 0===t.top&&void 0===t.bottom||void 0!==t.top&&(parseInt(u.getStyle(i,"top"))||0)===t.top||void 0!==t.bottom&&(parseInt(u.getStyle(i,"bottom"))||0)===t.bottom)){!0===t.stop?!0===z.isAttached&&z.detach():!1===z.isAttached&&z.attach();var e=t.position||C.css.position;C.position=e,i.style.position=e,i.style.left=!0===t.stop?C.css.left:z.positionLeft+"px",i.style.width="absolute"!==e?C.css.width:z.width,void 0===t.bottom?i.style.bottom=!0===t.stop?"":"auto":i.style.bottom=t.bottom+"px",void 0===t.top?i.style.top=!0===t.stop?"":"auto":i.style.top=t.top+"px",i.classList?i.classList.remove(p.stickyClass):i.className=i.className.replace(new RegExp("(^|\\b)"+p.stickyClass.split(" ").join("|")+"(\\b|$)","gi")," "),p.onStop&&p.onStop.call(i,Object.assign({},p))}}},z={el:$.createElement("div"),offsetLeft:null,positionLeft:null,width:null,isAttached:!1,init:function(){for(var t in z.el.className="sticky-spacer",C.css)z.el.style[t]=C.css[t];z.el.style["z-index"]="-1";var e=u.getStyle(i);z.offsetLeft=u.offset(i).left-(parseInt(e.marginLeft)||0),z.positionLeft=u.position(i).left,z.width=u.getStyle(i,"width")},attach:function(){s.insertBefore(z.el,i),z.isAttached=!0},detach:function(){z.el=s.removeChild(z.el),z.isAttached=!1}},n=function(){var t,e,o,n;C.css=(t=i,e=u.getCascadedStyle(t),o=u.getStyle(t),n={height:t.offsetHeight+"px",left:e.left,right:e.right,top:e.top,bottom:e.bottom,position:o.position,display:o.display,verticalAlign:o.verticalAlign,boxSizing:o.boxSizing,marginLeft:e.marginLeft,marginRight:e.marginRight,marginTop:e.marginTop,marginBottom:e.marginBottom,paddingLeft:e.paddingLeft,paddingRight:e.paddingRight},e.float&&(n.float=e.float||"none"),e.cssFloat&&(n.cssFloat=e.cssFloat||"none"),o.MozBoxSizing&&(n.MozBoxSizing=o.MozBoxSizing),n.width="auto"!==e.width?e.width:"border-box"===n.boxSizing||"border-box"===n.MozBoxSizing?t.offsetWidth+"px":o.width,n),z.init(),r=!(!p.stickTo||!("document"===p.stickTo||p.stickTo.nodeType&&9===p.stickTo.nodeType||"object"===_typeof(p.stickTo)&&p.stickTo instanceof("undefined"!=typeof HTMLDocument?HTMLDocument:Document))),l=p.stickTo?r?$:"string"==typeof p.stickTo?$.querySelector(p.stickTo):p.stickTo:s,E=(T=function(){var t=i.offsetHeight+(parseInt(C.css.marginTop)||0)+(parseInt(C.css.marginBottom)||0),e=(E||0)-t;return-1<=e&&e<=1?E:t})(),c=(L=function(){return r?Math.max($.documentElement.clientHeight,$.body.scrollHeight,$.documentElement.scrollHeight,$.body.offsetHeight,$.documentElement.offsetHeight):l.offsetHeight})(),d=r?0:u.offset(l).top,y=p.stickTo?r?0:u.offset(s).top:d,m=_.innerHeight,x=i.offsetTop-(parseInt(C.css.marginTop)||0),a=p.innerSticker?"string"==typeof p.innerSticker?$.querySelector(p.innerSticker):p.innerSticker:null,g=isNaN(p.top)&&-1<p.top.indexOf("%")?parseFloat(p.top)/100*m:p.top,h=isNaN(p.bottom)&&-1<p.bottom.indexOf("%")?parseFloat(p.bottom)/100*m:p.bottom,b=a?a.offsetTop:p.innerTop?p.innerTop:0,v=isNaN(p.bottomEnd)&&-1<p.bottomEnd.indexOf("%")?parseFloat(p.bottomEnd)/100*m:p.bottomEnd,S=d-g+b+x},N=_.pageYOffset||$.documentElement.scrollTop,H=0,R=function(){E=T(),c=L(),w=d+c-g-v,k=m<E;var t,e=_.pageYOffset||$.documentElement.scrollTop,o=u.offset(i).top,n=o-e;j=e<N?"up":"down",H=e-N,S<(N=e)?w+g+(k?h:0)-(p.followScroll&&k?0:g)<=e+E-b-(m-(S-b)<E-b&&p.followScroll&&0<(t=E-m-b)?t:0)?C.release({position:"absolute",bottom:y+s.offsetHeight-w-g}):k&&p.followScroll?"down"===j?n+E+h<=m+.9?C.stick({bottom:h}):"fixed"===C.position&&C.release({position:"absolute",top:o-g-S-H+b}):Math.ceil(n+b)<0&&"fixed"===C.position?C.release({position:"absolute",top:o-g-S+b-H}):e+g-b<=o&&C.stick({top:g-b}):C.stick({top:g-b}):C.release({stop:!0})},A=!1,B=!1,I=function(){A&&(u.event.unbind(_,"scroll",R),A=!1)},q=function(){null!==i.offsetParent&&"none"!==u.getStyle(i,"display")?(n(),c<=E?I():(R(),A||(u.event.bind(_,"scroll",R),A=!0))):I()},F=function(){i.style.position="",i.style.left="",i.style.top="",i.style.bottom="",i.style.width="",i.classList?i.classList.remove(p.stickyClass):i.className=i.className.replace(new RegExp("(^|\\b)"+p.stickyClass.split(" ").join("|")+"(\\b|$)","gi")," "),C.css={},!(C.position=null)===z.isAttached&&z.detach()},M=function(){F(),e(),t()?I():q()},D=function(){p.onBeforeResize&&p.onBeforeResize.call(i,Object.assign({},p)),M(),p.onResize&&p.onResize.call(i,Object.assign({},p))},P=p.resizeDebounce?u.debounce(D,p.resizeDebounce):D,W=function(){B&&(u.event.unbind(_,"resize",P),B=!1),I()},V=function(){B||(u.event.bind(_,"resize",P),B=!0),e(),t()?I():q()};this.options=function(t){return t?p[t]:Object.assign({},p)},this.refresh=M,this.update=function(t){O(t),f=Object.assign({},f,t||{}),M()},this.attach=V,this.detach=W,this.destroy=function(){W(),F()},this.triggerMethod=function(t,e){"function"==typeof o[t]&&o[t](e)},this.reinit=function(){Y("reinit","refresh","method"),M()},O(f),V(),u.event.bind(_,"load",M)};if(void 0!==_.jQuery){var n=_.jQuery,i="hcSticky";n.fn.extend({hcSticky:function(e,o){return this.length?"options"===e?n.data(this.get(0),i).options():this.each(function(){var t=n.data(this,i);t?t.triggerMethod(e,o):(t=new Q(this,e),n.data(this,i,t))}):this}})}return _.hcSticky=_.hcSticky||Q,Q}),function(c){"use strict";var t=c.hcSticky,f=c.document;"function"!=typeof Object.assign&&Object.defineProperty(Object,"assign",{value:function(t,e){if(null==t)throw new TypeError("Cannot convert undefined or null to object");for(var o=Object(t),n=1;n<arguments.length;n++){var i=arguments[n];if(null!=i)for(var s in i)Object.prototype.hasOwnProperty.call(i,s)&&(o[s]=i[s])}return o},writable:!0,configurable:!0}),Array.prototype.forEach||(Array.prototype.forEach=function(t){var e,o;if(null==this)throw new TypeError("this is null or not defined");var n=Object(this),i=n.length>>>0;if("function"!=typeof t)throw new TypeError(t+" is not a function");for(1<arguments.length&&(e=arguments[1]),o=0;o<i;){var s;o in n&&(s=n[o],t.call(e,s,o,n)),o++}});var e=function(){var t=f.documentElement,e=function(){};function n(t){var e=c.event;return e.target=e.target||e.srcElement||t,e}t.addEventListener?e=function(t,e,o){t.addEventListener(e,o,!1)}:t.attachEvent&&(e=function(e,t,o){e[t+o]=o.handleEvent?function(){var t=n(e);o.handleEvent.call(o,t)}:function(){var t=n(e);o.call(e,t)},e.attachEvent("on"+t,e[t+o])});var o=function(){};return t.removeEventListener?o=function(t,e,o){t.removeEventListener(e,o,!1)}:t.detachEvent&&(o=function(e,o,n){e.detachEvent("on"+o,e[o+n]);try{delete e[o+n]}catch(t){e[o+n]=void 0}}),{bind:e,unbind:o}}(),r=function(t,e){return c.getComputedStyle?e?f.defaultView.getComputedStyle(t,null).getPropertyValue(e):f.defaultView.getComputedStyle(t,null):t.currentStyle?e?t.currentStyle[e.replace(/-\w/g,function(t){return t.toUpperCase().replace("-","")})]:t.currentStyle:void 0},l=function(t){var e=t.getBoundingClientRect(),o=c.pageYOffset||f.documentElement.scrollTop,n=c.pageXOffset||f.documentElement.scrollLeft;return{top:e.top+o,left:e.left+n}};t.Helpers={isEmptyObject:function(t){for(var e in t)return!1;return!0},debounce:function(n,i,s){var r;return function(){var t=this,e=arguments,o=s&&!r;clearTimeout(r),r=setTimeout(function(){r=null,s||n.apply(t,e)},i),o&&n.apply(t,e)}},hasClass:function(t,e){return t.classList?t.classList.contains(e):new RegExp("(^| )"+e+"( |$)","gi").test(t.className)},offset:l,position:function(t){var e=t.offsetParent,o=l(e),n=l(t),i=r(e),s=r(t);return o.top+=parseInt(i.borderTopWidth)||0,o.left+=parseInt(i.borderLeftWidth)||0,{top:n.top-o.top-(parseInt(s.marginTop)||0),left:n.left-o.left-(parseInt(s.marginLeft)||0)}},getStyle:r,getCascadedStyle:function(t){var e,o=t.cloneNode(!0);o.style.display="none",Array.prototype.slice.call(o.querySelectorAll('input[type="radio"]')).forEach(function(t){t.removeAttribute("name")}),t.parentNode.insertBefore(o,t.nextSibling),o.currentStyle?e=o.currentStyle:c.getComputedStyle&&(e=f.defaultView.getComputedStyle(o,null));var n={};for(var i in e)!isNaN(i)||"string"!=typeof e[i]&&"number"!=typeof e[i]||(n[i]=e[i]);if(Object.keys(n).length<3)for(var s in n={},e)isNaN(s)||(n[e[s].replace(/-\w/g,function(t){return t.toUpperCase().replace("-","")})]=e.getPropertyValue(e[s]));if(n.margin||"auto"!==n.marginLeft?n.margin||n.marginLeft!==n.marginRight||n.marginLeft!==n.marginTop||n.marginLeft!==n.marginBottom||(n.margin=n.marginLeft):n.margin="auto",!n.margin&&"0px"===n.marginLeft&&"0px"===n.marginRight){var r=t.offsetLeft-t.parentNode.offsetLeft,l=r-(parseInt(n.left)||0)-(parseInt(n.right)||0),a=t.parentNode.offsetWidth-t.offsetWidth-r-(parseInt(n.right)||0)+(parseInt(n.left)||0)-l;0!==a&&1!==a||(n.margin="auto")}return o.parentNode.removeChild(o),o=null,n},event:e}}(window);`

//HCStickyJS - HC Sticky JS
func HCStickyJS() io.Reader {
	return strings.NewReader(hcstickyjs)
}

//HCOffCanvasNavJS - HC Offanvas Nav JS
func HCOffCanvasNavJS() io.Reader {
	return strings.NewReader(hcoffcanvasnavjs)
}

//HCOffCanvasNavCSS - HC OffCanvas Nav CSS
func HCOffCanvasNavCSS() io.Reader {
	return strings.NewReader(hcoffcanvasnavcss)
}