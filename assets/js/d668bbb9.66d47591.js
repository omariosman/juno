"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[9785],{3905:(e,t,r)=>{r.d(t,{Zo:()=>s,kt:()=>m});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var l=n.createContext({}),p=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):c(c({},t),e)),r},s=function(e){var t=p(e.components);return n.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},g=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,s=i(e,["components","mdxType","originalType","parentName"]),g=p(r),m=o,d=g["".concat(l,".").concat(m)]||g[m]||u[m]||a;return r?n.createElement(d,c(c({ref:t},s),{},{components:r})):n.createElement(d,c({ref:t},s))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,c=new Array(a);c[0]=g;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,c[1]=i;for(var p=2;p<a;p++)c[p]=r[p];return n.createElement.apply(null,c)}return n.createElement.apply(null,r)}g.displayName="MDXCreateElement"},8514:(e,t,r)=>{r.r(t),r.d(t,{contentTitle:()=>c,default:()=>s,frontMatter:()=>a,metadata:()=>i,toc:()=>l});var n=r(7462),o=(r(7294),r(3905));const a={title:"Coverage",description:"How to make coverage to the app"},c=void 0,i={unversionedId:"testing/coverage",id:"testing/coverage",title:"Coverage",description:"How to make coverage to the app",source:"@site/docs/testing/coverage.mdx",sourceDirName:"testing",slug:"/testing/coverage",permalink:"/docs/testing/coverage",editUrl:"https://github.com/NethermindEth/juno/tree/main/docs/docs/testing/coverage.mdx",tags:[],version:"current",frontMatter:{title:"Coverage",description:"How to make coverage to the app"},sidebar:"tutorialSidebar",previous:{title:"Testing",permalink:"/docs/category/testing"},next:{title:"Format",permalink:"/docs/testing/format"}},l=[{value:"Execution",id:"execution",children:[],level:2}],p={toc:l};function s(e){let{components:t,...r}=e;return(0,o.kt)("wrapper",(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("p",null,"We use ",(0,o.kt)("a",{parentName:"p",href:"https://about.codecov.io/"},"Codecov")," for coverage"),(0,o.kt)("h1",{id:"installation"},"Installation"),(0,o.kt)("p",null,"Install the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/DemerzelSolutions/courtney"},"courtney")," command-line tool to check configuration validity."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"make install-courtney\n")),(0,o.kt)("h2",{id:"execution"},"Execution"),(0,o.kt)("p",null,"For coverage generation:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"make codecov-test\n")),(0,o.kt)("p",null,"To display locally coverage tests:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"go tool cover -html=coverage/coverage.out\n")))}s.isMDXComponent=!0}}]);