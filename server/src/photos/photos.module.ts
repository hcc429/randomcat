import { Module } from '@nestjs/common';
import { PhotosService } from './photos.service';
import { PhotosController } from './photos.controller';

@Module({
  providers: [PhotosService],
  controllers: [PhotosController]
})
export class PhotosModule {}
