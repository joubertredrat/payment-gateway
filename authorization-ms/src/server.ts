import { Server, ServerCredentials } from '@grpc/grpc-js';

import { AuthorizationService } from './models/authorization';
import { HealthCheckResponse_ServingStatus, HealthService } from './models/health';
import { Authorization } from './server/Authorization';
import { Health, healthStatus } from './server/Health';
import { logger } from './utils';

const server = new Server({
  'grpc.max_receive_message_length': -1,
  'grpc.max_send_message_length': -1,
});

server.addService(AuthorizationService, new Authorization());
server.addService(HealthService, new Health());
server.bindAsync('0.0.0.0:50051', ServerCredentials.createInsecure(), (err: Error | null, bindPort: number) => {
  if (err) {
    throw err;
  }

  logger.info(`gRPC:Server:${bindPort}`, new Date().toLocaleString());
  server.start();
});

// Change service health status
healthStatus.set('helloworld.Greeter', HealthCheckResponse_ServingStatus.NOT_SERVING);
