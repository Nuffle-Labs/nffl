"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[509],{6898:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>d,frontMatter:()=>s,metadata:()=>a,toc:()=>p});var n=r(4848),o=r(8453);const s={sidebar_position:1},i="Registration",a={id:"operator/registration",title:"Registration",description:"Here we'll go step-by-step on how to opt-in into NFFL. It's a quick and",source:"@site/docs/operator/registration.md",sourceDirName:"operator",slug:"/operator/registration",permalink:"/operator/registration",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"sidebar",previous:{title:"Operator",permalink:"/category/operator"},next:{title:"Setup",permalink:"/operator/setup"}},l={},p=[{value:"Hardware Requirements",id:"hardware-requirements",level:2},{value:"Steps",id:"steps",level:2},{value:"Step 1: Complete EigenLayer Operator Registration",id:"step-1-complete-eigenlayer-operator-registration",level:3},{value:"Step 2: Install Docker",id:"step-2-install-docker",level:3},{value:"Step 3: Prepare Local NFFL files",id:"step-3-prepare-local-nffl-files",level:3},{value:"Step 4: Copy your EigenLayer operator keys to the setup directory",id:"step-4-copy-your-eigenlayer-operator-keys-to-the-setup-directory",level:3},{value:"Step 5: Update your <code>.env</code> file",id:"step-5-update-your-env-file",level:3},{value:"Step 6: Update your configuration files",id:"step-6-update-your-configuration-files",level:3},{value:"Step 6: Run the registration script",id:"step-6-run-the-registration-script",level:3}];function c(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.h1,{id:"registration",children:"Registration"}),"\n",(0,n.jsx)(t.p,{children:"Here we'll go step-by-step on how to opt-in into NFFL. It's a quick and\neasy process that will allow you to start contributing to the network once the\ntestnet starts functioning."}),"\n",(0,n.jsx)(t.h2,{id:"hardware-requirements",children:"Hardware Requirements"}),"\n",(0,n.jsxs)(t.p,{children:["The opt-in process is not hardware-intensive - you should be able to do it with\nlittle to no specific requirements. If you wish to use the same setup to run\nthe operator in the future, you can follow the hardware requirements on\n",(0,n.jsx)(t.a,{href:"./setup",children:"Setup"}),"."]}),"\n",(0,n.jsx)(t.h2,{id:"steps",children:"Steps"}),"\n",(0,n.jsx)(t.admonition,{type:"note",children:(0,n.jsx)(t.p,{children:"At this initial testnet stage, operators need to be whitelisted. If you are\ninterested and have not already been whitelisted, please contact the NFFL\nteam!"})}),"\n",(0,n.jsx)(t.h3,{id:"step-1-complete-eigenlayer-operator-registration",children:"Step 1: Complete EigenLayer Operator Registration"}),"\n",(0,n.jsxs)(t.p,{children:["Complete the EigenLayer CLI installation and registration ",(0,n.jsx)(t.a,{href:"https://docs.eigenlayer.xyz/operator-guides/operator-installation",children:"here"}),"."]}),"\n",(0,n.jsx)(t.h3,{id:"step-2-install-docker",children:"Step 2: Install Docker"}),"\n",(0,n.jsxs)(t.p,{children:["Install ",(0,n.jsx)(t.a,{href:"https://docs.docker.com/engine/install/ubuntu/",children:"Docker Engine on Linux"}),"."]}),"\n",(0,n.jsx)(t.h3,{id:"step-3-prepare-local-nffl-files",children:"Step 3: Prepare Local NFFL files"}),"\n",(0,n.jsx)(t.p,{children:"Clone the NFFL repository and execute the following."}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"git clone https://github.com/NethermindEth/near-sffl.git\ncd near-sffl/setup/plugin\ncp .env.example .env\n"})}),"\n",(0,n.jsx)(t.h3,{id:"step-4-copy-your-eigenlayer-operator-keys-to-the-setup-directory",children:"Step 4: Copy your EigenLayer operator keys to the setup directory"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"cp <path-to-your-operator-ecdsa-key> ./config/keys/ecdsa.json\ncp <path-to-your-operator-bls-key> ./config/keys/bls.json\n"})}),"\n",(0,n.jsxs)(t.h3,{id:"step-5-update-your-env-file",children:["Step 5: Update your ",(0,n.jsx)(t.code,{children:".env"})," file"]}),"\n",(0,n.jsxs)(t.p,{children:["You should have something similar to this in your ",(0,n.jsx)(t.code,{children:".env"}),":"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"# Operator BLS and ECDSA key passwords (from config/keys files)\nBLS_KEY_PASSWORD=fDUMDLmBROwlzzPXyIcy\nECDSA_KEY_PASSWORD=EnJuncq01CiVk9UbuBYl\n"})}),"\n",(0,n.jsxs)(t.p,{children:["For registering, set your EigenLayer ECDSA and BLS key passwords in the\n",(0,n.jsx)(t.code,{children:"ECDSA_KEY_PASSWORD"})," and ",(0,n.jsx)(t.code,{children:"BLS_KEY_PASSWORD"})," fields."]}),"\n",(0,n.jsx)(t.h3,{id:"step-6-update-your-configuration-files",children:"Step 6: Update your configuration files"}),"\n",(0,n.jsxs)(t.p,{children:["Now, in ",(0,n.jsx)(t.code,{children:"setup/plugin/config/operator.yaml"}),", set your ",(0,n.jsx)(t.code,{children:"operator_address"}),"\nand double-check the contract addresses."]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-yaml",children:"# Operator ECDSA address\noperator_address: 0xD5A0359da7B310917d7760385516B2426E86ab7f\n\n# AVS contract addresses\navs_registry_coordinator_address: 0x0069A298e68c09B047E5447b3b762E42114a99a2\noperator_state_retriever_address: 0x8D0b27Df027bc5C41855Da352Da4B5B2C406c1F0\n\n# AVS network RPCs\neth_rpc_url: https://ethereum-holesky-rpc.publicnode.com\neth_ws_url: wss://ethereum-holesky-rpc.publicnode.com\n\n# EigenLayer ECDSA and BLS private key paths\necdsa_private_key_store_path: /near-sffl/config/keys/ecdsa.json\nbls_private_key_store_path: /near-sffl/config/keys/bls.json\n"})}),"\n",(0,n.jsxs)(t.p,{children:["You'll need to refer to the ",(0,n.jsx)(t.a,{href:"./setup",children:"Setup"})," again before running the node for\nother important fields."]}),"\n",(0,n.jsx)(t.h3,{id:"step-6-run-the-registration-script",children:"Step 6: Run the registration script"}),"\n",(0,n.jsxs)(t.p,{children:["Now, simply run ",(0,n.jsx)(t.code,{children:"./register.sh"}),"! This will fetch our latest operator plugin\ncontainer and run it with the ",(0,n.jsx)(t.code,{children:"--operation-type opt-in"})," flag. It will\nopt-in your operator into NFFL."]})]})}function d(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(c,{...e})}):c(e)}},8453:(e,t,r)=>{r.d(t,{R:()=>i,x:()=>a});var n=r(6540);const o={},s=n.createContext(o);function i(e){const t=n.useContext(s);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function a(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:i(e.components),n.createElement(s.Provider,{value:t},e.children)}}}]);