import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";

const XSRF_HEADER = "X-Csrf-Token";

function getCsrf(): string {
    let maybeCsrf = document.querySelector('meta[name="csrf-token"]');
    if (maybeCsrf) {
        return maybeCsrf.getAttribute("content") || "";
    }

    return "";
}

const createAxiosClient = (baseUrl: string): AxiosInstance => {
    const options = {
        baseURL: baseUrl ?? window.location.host + "/api/",
    };

    let instance = axios.create(options);
    instance.defaults.headers.common[XSRF_HEADER] = getCsrf();

    return instance;
};

export class ApiClient {
    private client: AxiosInstance;

    constructor(client: AxiosInstance) {
        this.client = client;
    }

    get<T>(url: string, conf = {}) {
        return this.client
            .get<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    delete<T>(url: string, conf = {}) {
        return this.client
            .delete<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    head<T>(url: string, conf = {}) {
        return this.client
            .head<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    options<T>(url: string, conf = {}) {
        return this.client
            .options<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    post<T>(url: string, conf = {}) {
        return this.client
            .post<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    put<T>(url: string, conf = {}) {
        return this.client
            .put<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }

    patch<T>(url: string, conf = {}) {
        return this.client
            .patch<T>(url, conf)
            .then((response: AxiosResponse<T>) => Promise.resolve(response))
            .catch((error: Error | AxiosError) => Promise.reject(error));
    }
}

export default new ApiClient(
    createAxiosClient(`${window.location.origin}/api/v1`)
);
