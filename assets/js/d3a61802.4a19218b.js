"use strict";(self.webpackChunkjuno_docs=self.webpackChunkjuno_docs||[]).push([[8482],{3905:function(e,t,r){r.d(t,{Zo:function(){return u},kt:function(){return f}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var c=n.createContext({}),l=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,c=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),m=l(r),f=o,d=m["".concat(c,".").concat(f)]||m[f]||p[f]||i;return r?n.createElement(d,a(a({ref:t},u),{},{components:r})):n.createElement(d,a({ref:t},u))}));function f(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=m;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:o,a[1]=s;for(var l=2;l<i;l++)a[l]=r[l];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},3200:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return l},toc:function(){return u},default:function(){return m}});var n=r(7462),o=r(3366),i=(r(7294),r(3905)),a=["components"],s={title:"Juno Metrics"},c=void 0,l={unversionedId:"features/metrics",id:"features/metrics",title:"Juno Metrics",description:"Juno records metrics for the feeder gateway and Starknet Sync calls.",source:"@site/docs/features/metrics.mdx",sourceDirName:"features",slug:"/features/metrics",permalink:"/docs/features/metrics",editUrl:"https://github.com/NethermindEth/juno/tree/main/docs/docs/features/metrics.mdx",tags:[],version:"current",frontMatter:{title:"Juno Metrics"},sidebar:"tutorialSidebar",previous:{title:"Merkle Tries Details",permalink:"/docs/features/merkle-tree"},next:{title:"REST API",permalink:"/docs/features/rest"}},u=[],p={toc:u};function m(e){var t=e.components,r=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("p",null,"Juno records metrics for the ",(0,i.kt)("inlineCode",{parentName:"p"},"feeder")," gateway and ",(0,i.kt)("inlineCode",{parentName:"p"},"Starknet Sync")," calls."),(0,i.kt)("h1",{id:"viewing-the-metrics"},"Viewing the metrics"),(0,i.kt)("p",null,"To view the metrics, use this command:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"curl http://localhost:2048/metrics\n")),(0,i.kt)("h1",{id:"important-metrics"},"Important Metrics"),(0,i.kt)("p",null,"The important metrics that have been covered are as follows:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"no_of_requests")," - This metric provides the number of requests that have been sent to and received from the ",(0,i.kt)("inlineCode",{parentName:"p"},"feeder")," gateway. This metric not only shows the total number of requests but also depicts the number of requests for each type (for example - ",(0,i.kt)("inlineCode",{parentName:"p"},"get_contract_addresses")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"get_block"),"). In addition, this metric also stores the number of failures (if any) that occurred while making requests.")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"no_of_abi")," - This metric stores the number of ABI requests sent to and received from the feeder gateway. This also stores the number of failures (if any). ")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"count_starknet_sync")," - This metric stores the number of times Juno succeeds and fails in updating and committing blocks for syncing with Starknet.")),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"time_starknet_sync")," - This metric stores the total and average time needed for updating and committing blocks for syncing with Starknet."))))}m.isMDXComponent=!0}}]);