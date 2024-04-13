//
// Created by 14402 on 2023-06-04.
//
#include "simu8051_tools.h"
#include <stdio.h>
#include <stdlib.h>
#include "simu8051.h"

/**
 * 读取hex文件
 * @param buf   读取的缓冲区
 * @param width 字符宽度
 * @return int值
 */
static int read_hex(uint8_t *buf, int width) {
    int num = 0;

    for (int i = 0; i < width; i++) {
        char c = buf[i];

        if ((c >= '0') && (c <= '9')) {
            num = (num << 4) | (c - '0');
        }
        else if ((c >= 'a') && (c <= 'f')) {
            num = (num << 4) | (c - 'a') + 10;
        }
        else {
            num = (num << 4) | (c - 'A') + 10;
        }
    }

    return num;
}

/**
 * 加载hex文件
 * @param filename  文件名
 */
uint8_t *simu8051_load_hexfile(const char *filename) {
    char line_buf[HEX_LINE_SIZE];
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        printf("open file %s error\n", filename);
        return NULL;

    }
    // 8051 64kb
    uint8_t *code = (uint8_t *) malloc(1024 * 64);
    if (code == NULL) {
        printf("malloc fail\n");
        return NULL;
    }
    while (fgets(line_buf, sizeof(line_buf), file)) {
        uint8_t *c = line_buf;
//        hex 文件每行以:开始
        if (*c++ != ':') {
            fclose(file);
            free(code);
            return NULL;
        }
        // 获取本行数据长度
        uint8_t count = read_hex(c, 2);
        c += 2;
        // 获取数据起始地址
        uint16_t addr = read_hex(c, 4);
        c += 4;
        // 获取数据类型
        uint8_t type = read_hex(c, 2);
        c+=2;
        switch (type) {
            //数据类型
            case HEX_TYPE_DATA:
                for (uint8_t idx = 0; idx < count; idx++, c += 2) {
                    code[addr++] = read_hex(c, 2);
                }
                break;
                //结束
            case HEX_TYPE_EOF:
                fclose(file);
                return code;
            default:
                printf("unknow hex type %d\n", type);
                fclose(file);
                free(code);
                return NULL;
        }

    }
    fclose(file);
    return code;
}


void simu8051_dump_regs (void) {
    printf("\tR0=%2x, R1=%2x, R2=%2x, R3=%2x\n"
           "\tR4=%2x, R5=%2x, R6=%2x, R7=%2x,\n"
           "\ta=%2x,b=%2x,sp=%2x,dptr=%4x, \n"
           "\tpc=%4x, cyle=%4x, psw=%2x\n",
           simu8051_read(MEM_TYPE_IRAM, 0),
           simu8051_read(MEM_TYPE_IRAM, 1),
           simu8051_read(MEM_TYPE_IRAM, 2),
           simu8051_read(MEM_TYPE_IRAM, 3),
           simu8051_read(MEM_TYPE_IRAM, 4),
           simu8051_read(MEM_TYPE_IRAM, 5),
           simu8051_read(MEM_TYPE_IRAM, 6),
           simu8051_read(MEM_TYPE_IRAM, 7),
           simu8051_read(MEM_TYPE_SFR, SFR_ACC),
           simu8051_read(MEM_TYPE_SFR, SFR_B),
           simu8051_read(MEM_TYPE_SFR, SFR_SP),
           (simu8051_read(MEM_TYPE_SFR, SFR_DPH) << 8) | simu8051_read(MEM_TYPE_SFR, SFR_DPL),
           simu8051_pc(), simu8051_cycle(),
           simu8051_read(MEM_TYPE_SFR, SFR_PSW));
}