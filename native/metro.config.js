const {getDefaultConfig, mergeConfig} = require('@react-native/metro-config');
const fs = require('fs');
const path = require('path');
const exclusionList = require('metro-config/src/defaults/exclusionList');
const rnwPath = fs.realpathSync(path.resolve(require.resolve('react-native-windows/package.json'), '..'));
const MetroSymlinksResolver = require('@rnx-kit/metro-resolver-symlinks');
const {withNativeWind} = require('nativewind/metro');

//

/**
 * Metro configuration
 * https://facebook.github.io/metro/docs/configuration
 *
 * @type {import('metro-config').MetroConfig}
 */

// To be able to work with my ../../local_lib
const originalResolveRequest = MetroSymlinksResolver();

const config = mergeConfig(getDefaultConfig(__dirname), {
    //
    resolver: {
        blockList: exclusionList([
            // This stops "npx @react-native-community/cli run-windows" from causing the metro server to crash if its already running
            new RegExp(`${path.resolve(__dirname, 'windows').replace(/[/\\]/g, '/')}.*`),
            // This prevents "npx @react-native-community/cli run-windows" from hitting: EBUSY: resource busy or locked, open msbuild.ProjectImports.zip or other files produced by msbuild
            new RegExp(`${rnwPath}/build/.*`),
            new RegExp(`${rnwPath}/target/.*`),
            /.*\.ProjectImports\.zip/,
        ]),
        // This allows the framework dev mode
        resolveRequest: (context, moduleName, platform) => {
            // jotai "exports" messes with Metro
            if (moduleName.startsWith('jotai/')) {
                // for example: jotai/react, or jotai/vanilla
                return {
                    filePath: path.resolve(__dirname, `node_modules/${moduleName}.js`),
                    type: 'sourceFile',
                };
            }
            return originalResolveRequest(context, moduleName, platform);
        },
        // This also allows the framework dev mode
        extraNodeModules: {
            // because this allows to use our framework's code directly with code swapping
            goaldn: path.resolve(__dirname + '/../../goaldn'),
            emerald: path.resolve(__dirname + '/../../emeraldn'),

            // and this forces the code to get the statefull libs in our node modules, instead of the modules from our libs
            react: path.resolve(__dirname, 'node_modules/react'),
            'react-native': path.resolve(__dirname, 'node_modules/react-native'),
            'react-native-windows': path.resolve(__dirname, 'node_modules/react-native-windows'),
            '@react-navigation/native': path.resolve(__dirname, 'node_modules/@react-navigation/native'),
            'react-native-gesture-handler': path.resolve(__dirname, 'node_modules/react-native-gesture-handler'),
            'react-native-safe-area-context': path.resolve(__dirname, 'node_modules/react-native-safe-area-context'),
            'react-native-screens': path.resolve(__dirname, 'node_modules/react-native-screens'),
            'react-native-css-interop': path.resolve(__dirname, 'node_modules/react-native-css-interop'),
            'react-native-svg': path.resolve(__dirname, 'node_modules/react-native-svg'),
            'react-native-ble-manager': path.resolve(__dirname, 'node_modules/react-native-ble-manager'),
            'form-atoms': path.resolve(__dirname, 'node_modules/form-atoms'),
            'react-native-reanimated': path.resolve(__dirname, 'node_modules/react-native-reanimated'),
            'react-native-reanimated-carousel': path.resolve(__dirname, 'node_modules/react-native-reanimated-carousel'),
            jotai: path.resolve(__dirname, 'node_modules/jotai'),
        },
    },
    watchFolders: [path.resolve(__dirname + '/../../goaldn'), path.resolve(__dirname + '/../../emeraldn')],
    transformer: {
        getTransformOptions: async () => ({
            transform: {
                experimentalImportSupport: false,
                inlineRequires: true,
            },
        }),
    },
});

module.exports = withNativeWind(config, {input: './global.css'});
