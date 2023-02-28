import { Test, TestingModule } from '@nestjs/testing';
import { PhotosController } from './photos.controller';

describe('PhotosController', () => {
  let controller: PhotosController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [PhotosController],
    }).compile();

    controller = module.get<PhotosController>(PhotosController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
