import type { Config } from "tailwindcss";
import { aldesPreset } from "./src/vendor/emeraldn/styles/aldes-preset";
const plugin = require("tailwindcss/plugin");

const config = {
	presets: [require("nativewind/preset"), aldesPreset],
	content: [
		"./index.html",
		"./src/**/*.{ts,tsx}",
		"../../../open/emeraldn/lib/**/*.{ts,tsx}", // necessary for aldev swaps to work!
		"../../../open/goaldn/lib/**/*.{ts,tsx}", // necessary for aldev swaps to work!
	],
	theme: {
		// extend: {
		//     fontFamily: {
		//         sans: ['DMSans-Regular'],
		//     },
		// },
	},
	safelist: [
		...Array.from({ length: 100 }, (_, i) => `w-[${i + 1}%]`),
		// {
		//     pattern: /^w-\[\d+%\]$/, // Matches w-[25%], w-[100%]
		// },
		// 'w-[38%]',
	],
} satisfies Config;

export default config;
