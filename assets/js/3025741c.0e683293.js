"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[7513],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>d});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function a(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var u=r.createContext({}),c=function(e){var t=r.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},s=function(e){var t=c(e.components);return r.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,u=e.parentName,s=a(e,["components","mdxType","originalType","parentName"]),m=c(n),d=i,g=m["".concat(u,".").concat(d)]||m[d]||p[d]||o;return n?r.createElement(g,l(l({ref:t},s),{},{components:n})):r.createElement(g,l({ref:t},s))}));function d(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,l=new Array(o);l[0]=m;var a={};for(var u in t)hasOwnProperty.call(t,u)&&(a[u]=t[u]);a.originalType=e,a.mdxType="string"==typeof e?e:i,l[1]=a;for(var c=2;c<o;c++)l[c]=n[c];return r.createElement.apply(null,l)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},1491:(e,t,n)=>{n.r(t),n.d(t,{contentTitle:()=>l,default:()=>s,frontMatter:()=>o,metadata:()=>a,toc:()=>u});var r=n(7462),i=(n(7294),n(3905));const o={title:"Engineering Guidelines"},l="Patterns to follow",a={unversionedId:"contribution_guidelines/engineering-guidelines",id:"contribution_guidelines/engineering-guidelines",title:"Engineering Guidelines",description:"- Commit Conventions",source:"@site/docs/contribution_guidelines/engineering-guidelines.mdx",sourceDirName:"contribution_guidelines",slug:"/contribution_guidelines/engineering-guidelines",permalink:"/docs/contribution_guidelines/engineering-guidelines",editUrl:"https://github.com/NethermindEth/juno/tree/main/docs/docs/contribution_guidelines/engineering-guidelines.mdx",tags:[],version:"current",frontMatter:{title:"Engineering Guidelines"},sidebar:"tutorialSidebar",previous:{title:"Development Lifecycle",permalink:"/docs/contribution_guidelines/development-lifecycle"},next:{title:"Future Implementations",permalink:"/docs/category/future-implementations"}},u=[],c={toc:u};function s(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"patterns-to-follow"},"Patterns to follow"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://www.conventionalcommits.org/en/v1.0.0/"},"Commit Conventions"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://cbea.ms/git-commit/"},"How to Write a Git Commit Message")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://gojuno.xyz/contribution-guide"},"Contributor Guidelines")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://www.defmacro.org/2013/04/03/issue-etiquette.html"},"Github Issues Etiquette"))))}s.isMDXComponent=!0}}]);