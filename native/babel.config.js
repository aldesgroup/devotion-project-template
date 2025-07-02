module.exports = {
	presets: ["module:@react-native/babel-preset", "nativewind/babel"],
	plugins: [
		// ..other plugins if any
		[
			"module-resolver",
			{
				root: ["./src"],
				extensions: [".ts", ".tsx"],
				alias: {
					"~": "./src",
				},
			},
		],
	],
};
