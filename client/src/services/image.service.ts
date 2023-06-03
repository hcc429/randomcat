import ImageProps from "../interfaces/image";
import apiClient from "./api.client";

export default class ImageService {
  static async getPage(page: number, limit = 12): Promise<ImageProps[]> {
    return (await apiClient.get(`/images?p=${page}&limit=${limit}`)).data
      .images;
  }

  
}
