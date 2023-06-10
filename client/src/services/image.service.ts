import ImageProps from "../interfaces/image";
import Like from "../interfaces/Like";
import apiClient from "./api.client";

export default class ImageService {

  //TODO: error handling
  static async getPage(page: number, limit = 12): Promise<ImageProps[]> {
    return (await apiClient.get(`/images?p=${page}&limit=${limit}`)).data
      .images;
  }

  static async likeImages(likes: Like[]) {
    await apiClient.post(`/images/like`, { likes });
  }
}
