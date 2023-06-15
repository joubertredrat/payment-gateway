const { join } = require('path');
const { execSync } = require('child_process');
const { rimrafSync } = require('rimraf');

const PROTO_DIR = join(__dirname, '../../protos');
const MODEL_DIR = join(__dirname, '../src/models');
const PROTOC_PATH = join(__dirname, "../node_modules/grpc-tools/bin/protoc");
const PLUGIN_PATH = join(__dirname, "../node_modules/.bin/protoc-gen-ts_proto");

rimrafSync(`${MODEL_DIR}/*`, {
  glob: { ignore: `${MODEL_DIR}/tsconfig.json` },
});

const protoConfig = [
  `--plugin=${PLUGIN_PATH}`,

  // https://github.com/stephenh/ts-proto/blob/main/README.markdown
  "--ts_proto_opt=outputServices=grpc-js,env=node,useOptionals=messages,exportCommonSymbols=false,esModuleInterop=true",

  `--ts_proto_out=${MODEL_DIR}`,
  `--proto_path ${PROTO_DIR} ${PROTO_DIR}/*.proto`,
];

// https://github.com/stephenh/ts-proto#usage
execSync(`${PROTOC_PATH} ${protoConfig.join(" ")}`);
console.log(`> Proto models created: ${MODEL_DIR}`);
