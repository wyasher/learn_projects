#include <stdio.h>
#include <stdlib.h>
#include "simu8051.h"
#include "simu8051_tools.h"


#define REG_SP		0x1000
#define REG_A		0x1001
#define REG_B		0x1002
#define REG_PSW		0x1003
#define REG_PC		0x1004
#define REG_DPTR    0x1005
#define CYCLE   0x1006
#define REG_R0		0x2000
#define REG_R1		0x2001
#define REG_R2		0x2002
#define REG_R3		0x2003
#define REG_R4		0x2004
#define REG_R5		0x2005
#define REG_R6		0x2006
#define REG_R7		0x2007
#define REG_END     0x2FFF
static const char *file_names[] = {
        "code01.hex"
};

void check_result(void) {
    uint16_t addr = 0x600;
    do {
        uint8_t  addr0 = simu8051_read(MEM_TYPE_CODE, addr++);
        uint8_t  addr1 = simu8051_read(MEM_TYPE_CODE, addr++);
        uint8_t  data0 = simu8051_read(MEM_TYPE_CODE, addr++);
        uint8_t  data1 = simu8051_read(MEM_TYPE_CODE, addr++);
        uint16_t cmp_addr = (addr0 << 8) | addr1;
        uint16_t expect_data = (data0 << 8) | data1;

        uint16_t  real_data = 0;
            switch (cmp_addr) {
                case REG_SP:
                    real_data = simu8051_read(MEM_TYPE_SFR, SFR_SP);
                    break;
                case REG_A:
                    real_data = simu8051_read(MEM_TYPE_SFR, SFR_ACC);
                    break;
                case REG_B:
                    real_data = simu8051_read(MEM_TYPE_SFR, SFR_B);
                    break;
                case REG_PSW:
                    real_data = simu8051_read(MEM_TYPE_SFR, SFR_PSW);
                    break;
                case REG_PC:
                    real_data = simu8051_pc();
                    break;
                case REG_DPTR:
                    real_data = simu8051_read(MEM_TYPE_SFR, SFR_DPH) << 8;
                    real_data |= simu8051_read(MEM_TYPE_SFR, SFR_DPL);
                    break;
                    // cycle 不是一个寄存器
                case CYCLE:
                    real_data = simu8051_cycle();
                    break;
                case REG_R0:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 0);
                    break;
                case REG_R1:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 1);
                    break;
                case REG_R2:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 2);
                    break;
                case REG_R3:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 3);
                    break;
                case REG_R4:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 4);
                    break;
                case REG_R5:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 5);
                    break;
                case REG_R6:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 6);
                    break;
                case REG_R7:
                    real_data = simu8051_read(MEM_TYPE_IRAM, 7);
                    break;
                case REG_END:
                    return;
                default:
//                    其他数据
                    real_data = simu8051_read(MEM_TYPE_CODE, cmp_addr);
                    break;

            }
            if (real_data != expect_data) {
                printf("Error: addr: 0x%04X, expect: 0x%04X, real: 0x%04X\n", cmp_addr, expect_data, real_data);
                exit(1);
            }
    } while (1);
}

void test_memory(void) {
    static uint8_t data[MEM_XRAM_SIZE];
    for (int i = 0; i < 128; i++) {
        simu8051_write(MEM_TYPE_XRAM, i, 1);
        data[i] = simu8051_read(MEM_TYPE_XRAM, i);
    }
    for (int i = 128; i < 256; i++) {
        simu8051_write(MEM_TYPE_BIT, i, 1);
        data[i] = simu8051_read(MEM_TYPE_BIT, i);
    }

    for (int i = 0; i < MEM_IRAM_SIZE; i++) {
        simu8051_write(MEM_TYPE_IRAM, i, 1);
    }
    for (int i = 0; i < MEM_IRAM_SIZE; ++i) {
        data[i] = simu8051_read(MEM_TYPE_IRAM,1);
    }

    for (int i = 0; i < MEM_XRAM_SIZE; i++) {
        simu8051_write(MEM_TYPE_XRAM, i, 1);
        data[i] = simu8051_read(MEM_TYPE_XRAM,1);
    }
    simu8051_write(MEM_TYPE_SFR, SFR_ACC, 0x12);
    data[0] = simu8051_read(MEM_TYPE_SFR, SFR_ACC);

}
static void show_disa(instr_t* instr){
    const instr_info_t * info = instr->info;
    printf("c: %2d\n pc: %x \t",simu8051_cycle(),simu8051_pc());
    if (info->bytes == 1){
        printf("o: %x\t i: %s",instr->opcode,info->disa);
    } else if (info->bytes == 2){
        printf("o: %02x%02x\t i: %s",instr->opcode,instr->op0,info->disa);

    }else if (info->bytes == 3){
        printf("o: %02x%02x%02x\t i: %s",instr->opcode,instr->op0,instr->op1,info->disa);
    }
    printf("\n");
}
// 测试指令
void test_instr(void) {
    printf("begin test\n");
    for (int i = 0; i < sizeof(file_names) / sizeof(const char *); i++) {
        uint16_t pc;
        uint8_t *code;
        simu8051_reset();
        code = simu8051_load_hexfile(file_names[i]);
        if (code == NULL) {
            printf("load file %s failed\n", file_names[i]);
            exit(-1);
        }
        simu8051_load(code);
        do {
            instr_t instr;
            pc = simu8051_pc();
            simu8051_fetch_instr(&instr);
            show_disa(&instr);
            simu8051_exec(&instr);
            simu8051_dump_regs();
        }
            // 如果不是原地循环，就继续执行
        while (pc != simu8051_pc());
        check_result();
        free(code);
    }
    printf("end test\n");


}

int main() {
    simu8051_init();
//    test_memory();
    test_instr();

    return 0;
}
