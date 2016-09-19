/*****************************************************************************
   模块名      : constdef 
   文件名      : constdef.h
   相关文件    : 
   文件实现功能: constdef
   作者        : 顾爱华
   版本        : V1.0  Copyright(C) 2014-2015 HDDZ, All rights reserved.
-----------------------------------------------------------------------------
   修改记录:
   日  期      版本        修改人      修改内容
   2014/11/28  1.0         顾爱华        创建
   2016/01/30  1.1         顾爱华        重构
******************************************************************************/

#ifndef __CONSTDEF_H__
#define __CONSTDEF_H__

#ifdef __cplusplus
extern "C" {
#endif

#include "hdtype.h"

//! buf len
#define HDMAX_STR8_LEN          (u32)8
#define HDMAX_STR16_LEN         (u32)16
#define HDMAX_STR32_LEN         (u32)32
#define HDMAX_STR64_LEN         (u32)64
#define HDMAX_STR128_LEN        (u32)128
#define HDMAX_STR256_LEN        (u32)256
#define HDMAX_STR512_LEN        (u32)512
#define HDMAX_BUFF1K_LEN        (u32)1024
#define HDMAX_BUFF2K_LEN        (u32)(2*1024)
#define HDMAX_BUFF3K_LEN        (u32)(3*1024)
#define HDMAX_BUFF4K_LEN        (u32)(4*1024)
#define HDMAX_BUFF8K_LEN        (u32)(8*1024)

#define TPS_APPDELAY            (u32)30*1000	/*!< 30s */
#define TPS_45ms_TIMER          (u32)45
#define TPS_200ms_TIMER         (u32)200
#define TPS_500ms_TIMER         (u32)500
#define TPS_800ms_TIMER         (u32)800
#define TPS_1S_TIMER            (u32)1*1000
#define TPS_2S_TIMER            (u32)2*1000
#define TPS_3S_TIMER            (u32)3*1000
#define TPS_5S_TIMER            (u32)5*1000
#define TPS_8S_TIMER            (u32)8*1000
#define TPS_30S_TIMER           (u32)30*1000
#define TPS_45S_TIMER           (u32)45*1000
#define TPS_60S_TIMER           (u32)60*1000
#define TPS_2Min_TIMER          (u32)2*60*1000
#define TPS_5Min_TIMER          (u32)5*60*1000

#define PUSTATE_IDLE            (u8)0
#define PUSTATE_NORMAL          (u8)1

#define SYS_DEDAULT_PORT        (u16)2500
#define SYS_APP_NAME            (LPCSTR)"hdbox"
#define SYS_UPDATE_NAME         (LPCSTR)"hdbox_update"
#define SYS_UPDATE_TMP          (LPCSTR)"hdbox_update_tmp"
#define SYS_UART_NAME           (LPCSTR)"UARTGD"
#define SYS_PY_ABNORMAL         (LPCSTR)"ABNORMAL"

#define HD_INVALID_ID           (u8)0xFF
#define APP_MAIN                (u16)1
#define PUSRC_API               (u16)2
#define PUSRC_CFG               (u16)3
#define PUSRC_NET               (u16)4
#define PUSRC_SQL               (u16)5
#define PUSRC_MEDIA             (u16)6

//////////////////////////////////api////////////////////////////////////
#define BASE_LOG_PATH           (LPCSTR)"./log/"
#define BASE_LOG_INIT           (LPCSTR)"./log/init.dat"
#define BASE_LOG_RUN            (LPCSTR)"./log/run.dat"
#define BASE_LOG_EXT_TXT        (LPCSTR)".txt"
#define BASE_LOG_EXT_LOG        (LPCSTR)".log"
#define BASE_MAXLEN_INITLOG     (1*1024*1024)
#define BASE_MAXLEN_RUNLOG      (10*1024*1024)
#define BASE_MAXDAY_LOG         (u32)(30*24*3600)

#define BASE_IPMAC_CHK_INTERVAL     (s32)10
#define BASE_SYS_WATCHDOG           (LPCSTR)"/dev/watchdog"

#define BASE_NET_ETH0               (LPCSTR)"eth0"
#define BASE_NET_PPP0               (LPCSTR)"ppp0"
#define BASE_NET_WLAN0              (LPCSTR)"wlan0"

/////////////////////////////////cfg/////////////////////////////////////
#define CFG_PATH_CONFIG             (LPCSTR)"./config/hdcfg.ini"	/*!< 配置文件路径 */
#define CFG_PATH_MACREBOOT_FLAG     (LPCSTR)"./config/.macreboot.flag"

#define CFG_DEFAULT_IP              (LPCSTR)"224.1.1.1"
#define CFG_DEFAULT_PORT            (u16)7000

#define CFG_SNTPTM_IP               (LPCSTR)"202.120.2.101"
#define CFG_SNTPTM_PORT             (u16)123

#define CFG_DEFAULT_VER             (LPCSTR)"1.0"
//////////////////////////////////uart///////////////////////////////////
#define UART_DEV_RS485              (LPCSTR)"/dev/ttymxc1"
#define UART_DEV_RS232              (LPCSTR)"/dev/ttymxc2"
#define UART_MAXNUM_READ            (u8)6
#define UART_MAXNUM_PROTOCAL        (u8)16

/////////////////////////////////net////////////////////////////////////
#define SNTP_DEFAULT_DISTM          2208988800UL        /* 1970 - 1900 in seconds */

#define NET_MAX_PPPD_RSTNUM         (u8)3
#define NET_MAX_WCDMA_RSTNUM        (u8)3
#define NET_DEV_TTYUSB2             (LPCSTR)"/dev/ttyUSB2"

#define NET_SMS_SEND                (u8)0
#define NET_SMS_RECV                (u8)1

#define UDP_DEFAULT_PORT            (u16)50000
#define NET_SMS_DNS_CONF            (LPCSTR)"/etc/resolv.conf"

#define NET_MAX_SMSLEN              (u32)160*3
#define MAIL_RCV_MAX_NUM            (u8)5

/////////////////////////////////media///////////////////////////////////

/////////////////////////////////enum////////////////////////////////////

typedef enum
{
    NETMODE_DHCP     = 0,    //! DHCP
    NETMODE_STABLE,          //! 固定IP
}ENETWORKMODE;

typedef enum
{
    SMS_STAT_FAIL       = 0,
    SMS_STAT_DONE       = 1,
    SMS_STAT_WAIT       = 2,
}ESMSSTAT;

typedef enum
{
    MSG_TYPE_NONE    = 0,
    MSG_TYPE_TEST    = 1,
    MSG_TYPE_ALARM   = 2,
    MSG_TYPE_MAIL    = 3,
    MSG_TYPE_CURL    = 4,
}EMSGTYPE;

typedef enum
{
    CURL_MODE_GET       = 0,
    CURL_MODE_POST      = 1,
}ECURLMODE;

typedef enum
{
    ALARM_MODE_SMS      = 0,
    ALARM_MODE_MAIL     = 1,
    ALARM_MODE_CURL     = 2,
}EALARMMODE;

////////////////////////////////func////////////////////////////////////
#define __FUNC_LINE__ __FUNCTION__, __LINE__

#define API_ASSERT(var) \
    { \
        if (0 == var) \
        { \
            api_screen_printf("%s[%d] %s invalid pointer or value!\n", __FUNC_LINE__, #var); \
            return; \
        } \
    }

#define API_ASSERT_RET(var, ret) \
    { \
        if (0 == var) \
        { \
            api_screen_printf("%s[%d] %s invalid pointer or non value!\n", __FUNC_LINE__, #var); \
            return ret; \
        } \
    }

//! 适用于数组和结构体(指针不适用)
#define mzero(a) \
    { \
        memset(&(a), 0x00, sizeof(a)); \
    }

//! 获取数组长度
/*! \param  数组名
  !! 禁止对指针等使用
*/
#define arraynum(a)   (sizeof(a)/sizeof(a[0]))

#ifndef min
#define min(a, b) ((a) > (b) ? (b) : (a))
#endif

#ifndef max
#define max(a, b) ((a) < (b) ? (b) : (a))
#endif

// 将b拷入a中,a和b必须是以/0结尾的字符串(a,b均为数组)
#define mcopy( a, b ) \
{ \
	u32 len = min( sizeof( a ), strlen( b ) + 1 ); \
	memset( (a), 0, sizeof( a ) ); \
	memcpy( (a), (b), len ); \
	(a)[ len - 1 ] = '\0'; \
}

// 将b拷入a中,a和b必须指定长度(长度包括\0)
#define mcopyptr( a, b, alen, blen ) \
{ \
	u32 len = min( (alen), (blen) ); \
	memset( (a), 0, (alen) ); \
	memcpy( (a), (b), len ); \
	(a)[ len - 1 ] = '\0'; \
}

#define ip2str(ip, buf, buflen) \
{ \
	snprintf((char *)buf, buflen, "%lu.%lu.%lu.%lu", \
	(ip >>  0) & 0xFF, \
	(ip >>  8) & 0xFF, \
	(ip >>  16) & 0xFF, \
	(ip >>  24) & 0xFF); \
}

//! 安全释放内存
#define api_free(p) \
do {\
	if ((p) != NULL) \
	{\
		free(p); \
		(p) = NULL;\
	} \
} while(0)
    
//! 关闭文件指针
#define api_fclose(p) \
do {\
    if( NULL != p ) \
    { \
        fclose( p ); \
        p = NULL; \
    } \
} while(0)

#ifdef __cplusplus
}
#endif
#endif

