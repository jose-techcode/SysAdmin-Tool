package apiSysinfo

import (
	"net/http"

	"github.com/jose-techcode/SysAdmin-Tool/sysinfo"
	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()

	// All

	e.GET("/", func(c echo.Context) error {
		informations := sysinfo.Runsys()
		return c.String(http.StatusOK, informations)
	})

	// CPU

	e.GET("/cpu", func(c echo.Context) error {
		cpuInfo := sysinfo.GetCPUInfo()
		return c.String(http.StatusOK, cpuInfo)
	})

	// GPU

	e.GET("/gpu", func(c echo.Context) error {
		gpuInfo := sysinfo.GetGPUInfo()
		return c.String(http.StatusOK, gpuInfo)
	})

	// Motherboard

	e.GET("/motherboard", func(c echo.Context) error {
		motherboardInfo := sysinfo.GetMOTHERBOARDInfo()
		return c.String(http.StatusOK, motherboardInfo)
	})

	// BIOS

	e.GET("/bios", func(c echo.Context) error {
		biosInfo := sysinfo.GetBIOSInfo()
		return c.String(http.StatusOK, biosInfo)
	})

	// Memory

	e.GET("/mem", func(c echo.Context) error {
		memInfo := sysinfo.GetMEMInfo()
		return c.String(http.StatusOK, memInfo)
	})

	// Disk

	e.GET("/disk", func(c echo.Context) error {
		diskInfo := sysinfo.GetDISKInfo()
		return c.String(http.StatusOK, diskInfo)
	})

	// USB

	e.GET("/usb", func(c echo.Context) error {
		usbInfo := sysinfo.GetUSBInfo()
		return c.String(http.StatusOK, usbInfo)
	})

	// Host (Kernel & Operating System)

	e.GET("/host", func(c echo.Context) error {
		hostInfo := sysinfo.GetHOSTInfo()
		return c.String(http.StatusOK, hostInfo)
	})

	// Net

	e.GET("/net", func(c echo.Context) error {
		netInfo := sysinfo.GetNETInfo()
		return c.String(http.StatusOK, netInfo)
	})

	// Processes

	e.GET("/processes", func(c echo.Context) error {
		processesInfo := sysinfo.GetPROCESSESInfo()
		return c.String(http.StatusOK, processesInfo)
	})

	// Start the server

	e.Logger.Fatal(e.Start(":8000"))
}
