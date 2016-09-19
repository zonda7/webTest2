/*****************************************************************************
   Ä£¿éÃû      : hdstruct 
   ÎÄ¼þÃû      : hdstruct.h
   Ïà¹ØÎÄ¼þ    : 
   ÎÄ¼þÊµÏÖ¹¦ÄÜ: ¹«ÓÃstruct
   ×÷Õß        : ¹Ë°®»ª
   °æ±¾        : V1.0  Copyright(C) 2014-2015 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   ÐÞ¸Ä¼ÇÂ¼:
   ÈÕ  ÆÚ      °æ±¾        ÐÞ¸ÄÈË      ÐÞ¸ÄÄÚÈÝ
   2014/12/10  1.0         ¹Ë°®»ª        ´´½¨
   2016/01/30  1.1         ¹Ë°®»ª        ÖØ¹¹ 
******************************************************************************/

#ifndef __HDSTRUCT_H__
#define __HDSTRUCT_H__

#ifdef __cplusplus
extern "C" {
#endif

#include "constdef.h"

typedef struct __msg_header_t
{
	u16		wMsgId;						/*!< ÏûÏ¢ID */
	u16		wSrcAppId;					/*!< ÏûÏ¢À´Ô´APPID,·¢ËÍµÄÊ±ºò±íÊ¾Ä¿µÄAPP */
	u32		dwSrcNodeId;				/*!< ÏûÏ¢À´Ô´½áµãID,·¢ËÍµÄÊ±ºò±íÊ¾Ä¿µÄ½áµã */
	u32		dwHandle;					/*!< Á´Â·¾ä±ú */
	u16		wTransNo;					/*!< ÏûÏ¢´«ÊäID */
	u8		byChanId;					/*!< Í¨µÀID(Èç±à½âÂëÍ¨µÀ/´®¿Ú/²¢¿ÚµÈ) */
	u32		dwErrCode;					/*!< ´íÎóÂë(ÓÃÓÚÇëÇó·´À¡½á¹û) */
	u16		wMsgBodyLen;				/*!< ÏûÏ¢Ìå³¤¶È */
	u8		byByteOrder;				/*!< ×Ö½ÚÐò,Ä¬ÈÏÖ÷»úÐò(0:±¾»úÐò,1:ÍøÂçÐò) */
	u32		dwReserved;					/*!< ±£Áô×Ö¶Î,Ä¿Ç°ÓÃÓÚmpu±£´æ·¢ÆðÇëÇóµÄ×îÔ­Ê¼app */
}msg_header_t,*msg_header_pt;

typedef struct __msg_info_t
{
	msg_header_t tMsgHeader;
	u8	abyMsgBodyBuf[4*HDMAX_BUFF8K_LEN];
}msg_info_t, *msg_info_pt;

typedef struct __ipaddr_t
{
    u16 wPort;
    u32 dwIp;
}ipaddr_t, *ipaddr_pt;

typedef struct __time_info_t
{
	u16 wYear;
	u16 wMonth;
	u16 wDay;
	u16 wHour;
	u16 wMin;
	u16 wSec;
}time_info_t, *time_info_pt;

typedef struct __buf_data_t
{
    u32 dwCurTm;
    s32 nBufLen;
    s32 nRcvLen;
    u8 *achBuf;
}buf_data_t, *buf_data_pt;

//////////////////////////////////////////////////////////////////////////////
typedef struct __sntp_info_t
{
    u8 byEnb;
    u32 dwRate;
    u32 dwIp;
    u16 wPort;
    u8 byChangeMask;
}sntp_info_t, *sntp_info_pt;

typedef struct __sms_info_t
{
    u8 bySucc;          //! ESMSSTAT
    u8 byMode;          //! R/S
    u8 byType;          //! EMSGTYPE
    s8 achTime[HDMAX_STR32_LEN];
    s8 achPhone[HDMAX_STR32_LEN];
    s8 achCont[NET_MAX_SMSLEN];
}sms_info_t, *sms_info_pt;

typedef struct __sms_data_t
{
    u16 wCurNum;
    u16 wTotalNum;
    sms_info_t tSms[HDMAX_STR512_LEN];
}sms_data_t, *sms_data_pt;

typedef struct __sqlserver_info_t
{
    u16 wPort;
    s8 achAddr[HDMAX_STR32_LEN];
    s8 achUser[HDMAX_STR64_LEN];
    s8 achPasswd[HDMAX_STR128_LEN];
    s8 achDbName[HDMAX_STR128_LEN];
}sqlserver_info_t, *sqlserver_info_pt;

typedef struct __mail_c_t
{
    s8 achAddr[HDMAX_STR64_LEN];
}mail_c_t, *mail_c_pt;

typedef struct __mail_s_t
{
    s8 achUser[HDMAX_STR64_LEN];
    s8 achPasswd[HDMAX_STR64_LEN];
    s8 achServer[HDMAX_STR64_LEN];
}mail_s_t, *mail_s_pt;

typedef struct __mail_info_t
{
    u8 bySimAlarm;
    u8 byCvNum;
    u8 byCvTNum;
    mail_s_t tSv;
    mail_c_t tCvs[MAIL_RCV_MAX_NUM];
    u8 byChangeMask;
}mail_info_t, *mail_info_pt;

typedef struct __curl_info_t
{
    u8 byReqMode;  //! ECURLMODE
    s8 achUrl[HDMAX_STR128_LEN];
    s8 achPri[HDMAX_STR128_LEN];
    s8 achMid[HDMAX_STR128_LEN];
    s8 achSuf[HDMAX_STR128_LEN];
    u8 byChangeMask;
}curl_info_t, *curl_info_pt;

typedef struct __stat_show_t
{
    u8 bSysGuard;
    u8 bIsCard;
    u8 byRssi;
    u8 bInternet;
    u32 dwUptime;
    u8 achE0Mac[6];
    s8 achImei[HDMAX_STR16_LEN];
}stat_show_t, *stat_show_pt;

typedef struct __stat_option_t
{
    u8 bySysGuard;
    u8 byEnSqlSv;
    u8 byAlarmMode;     //! EALARMMODE
    u8 byChangeMask;
}stat_option_t, *stat_option_pt;

typedef struct __md5_info_t
{
    u8 achLow[HDMAX_STR16_LEN];
    u8 achUp[HDMAX_STR16_LEN];
}md5_info_t, *md5_info_pt;

//////////////////////////////////////////////////////////////////////////////
#pragma pack(push, 1)

typedef struct __ipv4_t
{
    unsigned long dwLocalIp;
    unsigned long dwNetMask;
    unsigned long dwGateWay;
    unsigned long dwFirstDns;
    unsigned long dwSecDns;
}ipv4_t, *ipv4_pt;

typedef struct __ipv4_info_t
{
    u8 byNetMode;       //! ENETWORKMODE
    ipv4_t tIpv4;
    u8 bySwitch;
    u8 byChangeMask;
}ipv4_info_t, *ipv4_info_pt;

typedef struct __lan_msg_t
{
    u16 wMsgId;
    u16 wErrCode;
    u16 wMsgBodyLen;
    u8  abyMsgBodyBuf[HDMAX_BUFF1K_LEN];
}lan_msg_t, *lan_msg_pt;

typedef struct __net_conflict_info_t
{
    u32 dwIpAddr;
    u8 macAddr[6];
}net_conflict_info_t, *net_conflict_info_pt;

#pragma pack(pop)

#ifdef __cplusplus
}
#endif

#endif

