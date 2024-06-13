// @ts-nocheck
/* eslint-disable */
/**
 * Masters way API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  SchemasUserPopulatedResponse,
} from '../models/index';
import {
    SchemasUserPopulatedResponseFromJSON,
    SchemasUserPopulatedResponseToJSON,
} from '../models/index';

/**
 * 
 */
export class GetCurrentAuthorizedUserApi extends runtime.BaseAPI {

    /**
     * Get current authorized user
     */
    async getCurrentAuthorizedUserRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasUserPopulatedResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/auth/current`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasUserPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Get current authorized user
     */
    async getCurrentAuthorizedUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasUserPopulatedResponse> {
        const response = await this.getCurrentAuthorizedUserRaw(initOverrides);
        return await response.value();
    }

}
