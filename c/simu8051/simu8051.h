//
// Created by 14402 on 2023-06-04.
//

#ifndef SIMU8051_H
#define SIMU8051_H
#include <stdint.h>
#define SFR_ACC 0xE0
#define SFR_SP 0x81
#define SFR_B 0xF0
#define SFR_DPH 0x83
#define SFR_DPL 0x82
#define SFR_PSW 0xD0
#define MEM_BIT_SFR_START 0x80
#define MEM_BIT_IRAM_START 0x20
#define MEM_CODE_SIZE 65536
#define MEM_IRAM_SIZE 128
#define MEM_SFR_SIZE 128
#define MEM_XRAM_SIZE 65536
#define SFR_ACC 0xE0
// 存储器类型
typedef enum _mem_type_t{
    MEM_TYPE_IRAM,
    MEM_TYPE_XRAM,
    MEM_TYPE_CODE,
    MEM_TYPE_BIT,
    MEM_TYPE_SFR,
} mem_type_t;
// 程序存储器结构体
typedef struct _memory_t{
    uint8_t  *code;
    uint8_t  iram[MEM_IRAM_SIZE];
    uint8_t  sfr[MEM_SFR_SIZE];
    // 外部
    uint8_t  xram[MEM_XRAM_SIZE];
} memory_t;


// 虚拟机描述结构体

typedef struct _simu8051_t{

    memory_t mem;
    // 程序计数器
    uint16_t  pc;
    uint32_t  cycles;
} simu8051_t;
struct _instr_info_t;
// 指令结构体
typedef struct _instr_t{
    uint8_t  opcode;// 操作码
    uint8_t  op0;
    uint8_t  op1;
    const struct _instr_info_t *info;
} instr_t;
// 操作数类型
typedef enum _op_type_t{
    OP_TYPE_NONE,
}op_type_t;
// 指令信息结构体
typedef struct _instr_info_t{
    uint8_t bytes;
    uint8_t cycles;
    op_type_t op0_type;
    op_type_t op1_type;
    void (*exec) (instr_t *instr);
    const char *disa;
} instr_info_t;

uint8_t simu8051_read(mem_type_t type, uint16_t addr);
void simu8051_write(mem_type_t type, uint16_t addr, uint8_t data);
void simu8051_reset(void );
void simu8051_init(void );
uint32_t simu8051_cycle(void );

// 获取程序
void simu8051_load(uint8_t *code);
// 获取程序计数器运行到哪里
uint16_t simu8051_pc(void);
// 获取指令
void simu8051_fetch_instr(instr_t *instr);
// 执行指令
void simu8051_exec(instr_t *instr);


#endif //SIMU8051_H
