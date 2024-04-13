//
// Created by 14402 on 2023-06-04.
//

#ifndef SIMU8051_TOOLS_H
#define SIMU8051_TOOLS_H
#include <stdint.h>
#define HEX_LINE_SIZE 1024
#define HEX_TYPE_DATA 0x00
#define HEX_TYPE_EOF 0x01
uint8_t* simu8051_load_hexfile(const char* filename);
void simu8051_dump_regs(void);
#endif //SIMU8051_TOOLS_H