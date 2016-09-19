/*****************************************************************************
   模块名      : HD type 
   文件名      : hdtype.h
   相关文件    : 
   文件实现功能: HD宏定义
   作者        : 顾爱华
   版本        : V1.0  Copyright(C) 2014-2015 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   修改记录:
   日  期      版本        修改人      修改内容
   2014/11/28  1.0         顾爱华        创建
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

