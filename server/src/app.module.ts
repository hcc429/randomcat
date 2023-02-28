import { Module } from '@nestjs/common';
import { ApiModule } from './api/api.module';
import { PhotosModule } from './photos/photos.module';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConfigModule, ConfigService } from '@nestjs/config';
import config from './config';
@Module({
  imports: [
    ApiModule,
    PhotosModule,
    ConfigModule.forRoot({
      envFilePath: ['.env', '.env.development'],
      isGlobal: true,
      load: [config],
    }),
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => configService.get('db'),
      imports: [ConfigModule],
    }),
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
