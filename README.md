# protoc-gen-js

## 简介

 读取protobuffer文件，自动生成api请求。

## 模板

```javaScripts
export function getAccountList(params: ObjTy) {
  return request({
    url: '/ad-manager/v1/browser/authorization/getAccountList',
    method: 'post',
    data: params,
    isJson:true,
    isParams:false
  })
}
```

