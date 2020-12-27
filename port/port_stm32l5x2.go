// +build stm32l5x2

package port

import "device/stm32"

var (
	GPIOA = &Port{
		GPIO_Type:           stm32.GPIOA,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOAEN,
	}
	GPIOB = &Port{
		GPIO_Type:           stm32.GPIOB,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOBEN,
	}
	GPIOC = &Port{
		GPIO_Type:           stm32.GPIOC,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOCEN,
	}
	GPIOD = &Port{
		GPIO_Type:           stm32.GPIOD,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIODEN,
	}
	GPIOE = &Port{
		GPIO_Type:           stm32.GPIOE,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOEEN,
	}
	GPIOF = &Port{
		GPIO_Type:           stm32.GPIOF,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOFEN,
	}
	GPIOG = &Port{
		GPIO_Type:           stm32.GPIOG,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOGEN,
		PowerEnableRegister: &stm32.PWR.CR2,
		PowerEnableFlag:     stm32.PWR_CR2_IOSV,
	}
	GPIOH = &Port{
		GPIO_Type:           stm32.GPIOH,
		ClockEnableRegister: &stm32.RCC.AHB2ENR,
		ClockEnableFlag:     stm32.RCC_AHB2ENR_GPIOHEN,
	}
)
