package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var macVendors map[string]string

func init() {
	macVendors = make(map[string]string)
	if err := loadMacVendorsFromURL("https://www.wireshark.org/download/automated/data/manuf"); err != nil {
		panic(fmt.Sprintf("Error loading the file: %v", err))
	}
}

func loadMacVendorsFromURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download the file: status code %d", resp.StatusCode)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		macPrefix := parts[0]
		vendorName := strings.Join(parts[2:], " ")
		macVendors[macPrefix] = vendorName
	}
	return scanner.Err()
}

func LookupVendor(c *fiber.Ctx) error {
	mac := c.Params("mac")
	vendor := lookupVendor(mac)
	if vendor == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Manufacturer not found"})
	}
	return c.JSON(fiber.Map{"mac": mac, "vendor": vendor})
}

func lookupVendor(mac string) string {
	mac = strings.ToUpper(mac)
	mac = strings.ReplaceAll(mac, "-", ":")
	mac = strings.ReplaceAll(mac, ".", ":")
	macParts := strings.Split(mac, ":")

	if len(macParts) < 3 {
		return ""
	}

	macPrefix := strings.Join(macParts[:3], ":")
	if vendor, exists := macVendors[macPrefix]; exists {
		return vendor
	}
	return ""
}
