/*****************************************************************************
   ģ����      : 
   �ļ���      : py_api.h
   ����ļ�    : 
   �ļ�ʵ�ֹ���: py_api
   ����        : �˰���
   �汾        : V1.0  Copyright(C) 2014-2016 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   �޸ļ�¼:
   ��  ��      �汾        �޸���      �޸�����
   2016/08/12  1.0         �˰���        ����
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

//! ��ʼ��opcIp
void PyInit();

struct ipv4_t* py_get_ipv4();

int py_set_ipv4(const struct ipv4_t* ptIpv4);

#ifdef __cplusplus
}
#endif
#endif

