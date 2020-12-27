// +build stm32f7x2

package port

import "device/stm32"

var (
	GPIOA = &Port{
		GPIO_Type:           stm32.GPIOA,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOAEN,
	}
	GPIOB = &Port{
		GPIO_Type:           stm32.GPIOB,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOBEN,
	}
	GPIOC = &Port{
		GPIO_Type:           stm32.GPIOC,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOCEN,
	}
	GPIOD = &Port{
		GPIO_Type:           stm32.GPIOD,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIODEN,
	}
	GPIOE = &Port{
		GPIO_Type:           stm32.GPIOE,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOEEN,
	}
	GPIOF = &Port{
		GPIO_Type:           stm32.GPIOF,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOFEN,
	}
	GPIOG = &Port{
		GPIO_Type:           stm32.GPIOG,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOGEN,
	}
	GPIOH = &Port{
		GPIO_Type:           stm32.GPIOH,
		ClockEnableRegister: &stm32.RCC.AHB1ENR,
		ClockEnableFlag:     stm32.RCC_AHB1ENR_GPIOHEN,
	}
)
