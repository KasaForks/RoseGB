package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"flag"
)

func main() {

	// init cpu
	cpu := CPU{}

	// check skip bootrom flag
	if checkBIOS() == false {
		cpu.bus.cartridge.loadBootROM()
	} else {
		cpu.skipBootROM()
	}

	title := "RoseGB - " + string(cpu.bus.cartridge.header.title[:])
	rl.InitWindow(160, 144, title)

	rl.SetTargetFPS(60)

	


	// mainloop which is executed 60 times per second
	for !rl.WindowShouldClose() {
		for i := 0; i < 70224; i++ {
			cpu.tick()
		}
		// rl.BeginDrawing()
		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawRectangle(1, 1, 1, 1, rl.Red)
		

		// rl.EndDrawing()
	}

	rl.CloseWindow()
}

func checkBIOS() bool {
	boolPtr := flag.Bool("skip-bootrom", false, "user specifies whether the bootrom should be loaded or not")
	flag.Parse()
	return *boolPtr
}

func (cpu *CPU) skipBootROM() {
	cpu.A = 0x01
	cpu.B = 0x00
	cpu.C = 0x13
	cpu.D = 0x00
	cpu.E = 0xD8
	cpu.F = 0xB0
	cpu.H = 0x01
	cpu.L = 0x4D
	cpu.SP = 0xFFFE
	cpu.PC = 0x100
	cpu.bus.write(0xFF05, 0x00)
	cpu.bus.write(0xFF06, 0x00)
  	cpu.bus.write(0xFF07, 0x00)
  	cpu.bus.write(0xFF10, 0x80)
  	cpu.bus.write(0xFF11, 0xBF)
 	cpu.bus.write(0xFF12, 0xF3)
 	cpu.bus.write(0xFF14, 0xBF)
 	cpu.bus.write(0xFF16, 0x3F)
 	cpu.bus.write(0xFF17, 0x00)
 	cpu.bus.write(0xFF19, 0xBF)
 	cpu.bus.write(0xFF1A, 0x7F)
 	cpu.bus.write(0xFF1B, 0xFF)
 	cpu.bus.write(0xFF1C, 0x9F)
 	cpu.bus.write(0xFF1E, 0xBF)
 	cpu.bus.write(0xFF20, 0xFF)
 	cpu.bus.write(0xFF21, 0x00)
 	cpu.bus.write(0xFF22, 0x00)
 	cpu.bus.write(0xFF23, 0xBF)
 	cpu.bus.write(0xFF24, 0x77)
 	cpu.bus.write(0xFF25, 0xF3)
 	cpu.bus.write(0xFF26, 0xF1) // F0 on super gameboy
 	cpu.bus.write(0xFF40, 0x91)
  	cpu.bus.write(0xFF42, 0x00)
  	cpu.bus.write(0xFF43, 0x00)
  	cpu.bus.write(0xFF45, 0x00)
  	cpu.bus.write(0xFF47, 0xFC)
  	cpu.bus.write(0xFF48, 0xFF)
  	cpu.bus.write(0xFF49, 0xFF)
  	cpu.bus.write(0xFF4A, 0x00)
  	cpu.bus.write(0xFF4B, 0x00)
  	cpu.bus.write(0xFFFF, 0x00)
  	cpu.bus.cartridge.loadCartridge()
}