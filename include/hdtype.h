/*****************************************************************************
   ģ����      : HD type 
   �ļ���      : hdtype.h
   ����ļ�    : 
   �ļ�ʵ�ֹ���: HD�궨��
   ����        : �˰���
   �汾        : V1.0  Copyright(C) 2014-2015 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   �޸ļ�¼:
   ��  ��      �汾        �޸���      �޸�����
   2014/11/28  1.0         �˰���        ����
******************************************************************************/

#ifndef __HDTYPE_H__
#define __HDTYPE_H__

#ifdef __cplusplus
extern "C" {
#endif

#define _LINUX_

typedef char s8;
typedef short s16;
typedef int s32, BOOL32;
typedef long long s64;

typedef unsigned char u8;
typedef unsigned short u16;
typedef unsigned long u32;
typedef unsigned long long u64;

#ifndef BOOL
#define BOOL int
#endif

#ifndef LPSTR
#define LPSTR char * 
#endif

#ifndef LPCSTR
#define LPCSTR const char *
#endif

#ifndef TRUE
#define TRUE 1
#endif

#ifndef FALSE
#define FALSE 0
#endif

#ifdef __cplusplus
}
#endif
#endif

