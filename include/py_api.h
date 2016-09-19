/*****************************************************************************
   模块名      : 
   文件名      : py_api.h
   相关文件    : 
   文件实现功能: py_api
   作者        : 顾爱华
   版本        : V1.0  Copyright(C) 2014-2016 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   修改记录:
   日  期      版本        修改人      修改内容
   2016/08/12  1.0         顾爱华        创建
******************************************************************************/

#ifndef _PY_API_H_
#define _PY_API_H_

#ifdef __cplusplus
extern "C" {
#endif

#include "hdtype.h"
#include "constdef.h"
//#include "hdstruct.h"
struct ipv4_t
{
    unsigned long dwLocalIp;
    unsigned long dwNetMask;
    unsigned long dwGateWay;
    unsigned long dwFirstDns;
    unsigned long dwSecDns;
};

//! 初始化opcIp
void PyInit();

struct ipv4_t* py_get_ipv4();

int py_set_ipv4(const struct ipv4_t* ptIpv4);

#ifdef __cplusplus
}
#endif
#endif

