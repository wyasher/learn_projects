//
// Created by 14402 on 2023-06-04.
//
#include "simu8051.h"
#include <string.h>
#include <stdio.h>

static simu8051_t simu;

static void do_nop(instr_t *instr) {
            simu.pc += instr->info->bytes;
            simu.cycles+=instr->info->cycles;
}
static void do_sjmp(instr_t *instr) {
            simu.pc += instr->info->bytes;
            // 向前跳还是向后跳转，所有需要把op0转换成有符号的int8_t
            simu.pc += (int8_t) instr->op0;
            simu.cycles+=instr->info->cycles;
}
// 指令表
static const instr_info_t instr_table[] = {
        [0x00] = {1,1,OP_TYPE_NONE,OP_TYPE_NONE,do_nop,"NOP"},
        [0x80] ={2,2,OP_TYPE_NONE,OP_TYPE_NONE,do_sjmp,"SJMP re"}
};

static uint8_t read_bit(uint8_t addr) {
    uint8_t bit_idx = addr % 8;
    if (addr < MEM_BIT_SFR_START) {
        uint8_t byte_index = addr / 8 + MEM_BIT_IRAM_START;
        return simu.mem.iram[byte_index] & (1 << bit_idx) ? 1 : 0;
    } else {
        uint8_t byte_index = (addr - MEM_BIT_SFR_START) / 8 * 8;
        return simu.mem.iram[byte_index] & (1 << bit_idx) ? 1 : 0;
    }
}

static uint8_t write_bit(uint8_t addr, int bit) {
    uint8_t bit_idx = addr % 8;
    bit &= 0x1;
    if (addr < MEM_BIT_SFR_START) {
        uint8_t byte_index = addr / 8 + MEM_BIT_IRAM_START;
        simu.mem.iram[byte_index] &= ~(1 << bit_idx);
        simu.mem.iram[byte_index] |= bit ? (1 << bit_idx) : 0;
    } else {
        uint8_t byte_index = (addr - MEM_BIT_SFR_START) / 8 * 8;
        simu.mem.sfr[byte_index] &= ~(1 << bit_idx);
        simu.mem.sfr[byte_index] |= bit ? (1 << bit_idx) : 0;
    }
}

static uint8_t read_sfr(uint8_t addr) {
    switch (addr) {
        default:
            addr -= MEM_BIT_SFR_START;
            return addr < MEM_TYPE_SFR ? simu.mem.sfr[addr] : 0;
    }
}

static void write_sfr(uint8_t addr, uint8_t data) {
    switch (addr) {
        default:
            addr -= MEM_BIT_SFR_START;
            if (addr < MEM_TYPE_SFR) {
                simu.mem.sfr[addr] = data;
            }
            break;
    }
}
// 重置虚拟机
void simu8051_reset() {
    simu.pc = 0;
    simu.cycles = 0;
    write_sfr(SFR_SP, 0x07);
}

//初始化虚拟机
void simu8051_init() {
    simu.mem.code = (uint8_t *) 0;
    memset(simu.mem.xram, 0, MEM_XRAM_SIZE);
    memset(simu.mem.iram, 0, MEM_IRAM_SIZE);
    memset(simu.mem.sfr, 0, MEM_SFR_SIZE);
    simu8051_reset();

}



// 读取内部ram
static uint8_t read_iram(uint8_t addr) {
    uint8_t data;
//    高128个字节是sfr

    if (addr >= MEM_IRAM_SIZE) {
        data = read_sfr(addr);
    } else {
        data = simu.mem.iram[addr];
    }
    return data;
}

static void write_iram(uint8_t addr, uint8_t data) {
//    高128个字节是sfr
    if (addr >= MEM_IRAM_SIZE) {
        write_sfr(addr, data);
    } else {
        simu.mem.iram[addr] = data;
    }
}

// xram uint16_t 最大 64kb
static uint8_t read_xram(uint16_t addr) {
    return simu.mem.xram[addr];
}

static void write_xram(uint16_t addr, uint8_t data) {
    simu.mem.xram[addr] = data;
}

// code
static uint8_t read_code(uint16_t addr) {
    return simu.mem.code[addr];
}

static void write_code(uint16_t addr, uint8_t data) {
    simu.mem.code[addr] = data;
}


uint8_t simu8051_read(mem_type_t type, uint16_t addr) {
    switch (type) {
        case MEM_TYPE_CODE:
            return read_code(addr);
        case MEM_TYPE_IRAM:
            return read_iram((uint8_t) addr);
        case MEM_TYPE_SFR:
            return read_sfr((uint8_t) addr);
        case MEM_TYPE_XRAM:
            return read_xram(addr);
        case MEM_TYPE_BIT:
            return read_bit((uint8_t) addr);
        default:
            return 0;

    }
}

void simu8051_write(mem_type_t type, uint16_t addr, uint8_t data) {
    switch (type) {
        case MEM_TYPE_CODE:
            write_code(addr, data);
            break;
        case MEM_TYPE_IRAM:
            write_iram((uint8_t) addr, data);
            break;
        case MEM_TYPE_SFR:
            write_sfr((uint8_t) addr, data);
            break;
        case MEM_TYPE_XRAM:
            write_xram(addr, data);
            break;
        case MEM_TYPE_BIT:
            write_bit(addr, data);
            break;
        default:
            break;
    }
}

void simu8051_load(uint8_t *code) {
    simu.mem.code = code;
}

// 获取程序计数器运行到哪里
uint16_t simu8051_pc(void) {
    return simu.pc;
}

// 获取指令
void simu8051_fetch_instr(instr_t *instr) {
    instr->opcode = simu.mem.code[simu.pc];
    instr->info = instr_table + instr->opcode;
    // 暂时不判断有两个字节指令 还是一个字节
    instr->op0 = simu.mem.code[simu.pc + 1];
    instr->op1 = simu.mem.code[simu.pc + 2];
}

// 执行指令
void simu8051_exec(instr_t *instr) {
    instr->info->exec(instr);
}
uint32_t simu8051_cycle(void) {
    return simu.cycles;
}

