// @ts-nocheck
/* eslint-disable */
/**
 * Masters way general API
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
  SchemasCreatePlanPayload,
  SchemasPlanPopulatedResponse,
  SchemasUpdatePlanPayload,
} from '../models/index';
import {
    SchemasCreatePlanPayloadFromJSON,
    SchemasCreatePlanPayloadToJSON,
    SchemasPlanPopulatedResponseFromJSON,
    SchemasPlanPopulatedResponseToJSON,
    SchemasUpdatePlanPayloadFromJSON,
    SchemasUpdatePlanPayloadToJSON,
} from '../models/index';

export interface CreatePlanRequest {
    request: SchemasCreatePlanPayload;
}

export interface DeletePlanRequest {
    planId: string;
}

export interface UpdatePlanRequest {
    planId: string;
    request: SchemasUpdatePlanPayload;
}

/**
 * 
 */
export class PlanApi extends runtime.BaseAPI {

    /**
     * Create a new plan
     */
    async createPlanRaw(requestParameters: CreatePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasPlanPopulatedResponse>> {
        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling createPlan.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/plans`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasCreatePlanPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasPlanPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Create a new plan
     */
    async createPlan(requestParameters: CreatePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasPlanPopulatedResponse> {
        const response = await this.createPlanRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete plan by UUID
     */
    async deletePlanRaw(requestParameters: DeletePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.planId === null || requestParameters.planId === undefined) {
            throw new runtime.RequiredError('planId','Required parameter requestParameters.planId was null or undefined when calling deletePlan.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/plans/{planId}`.replace(`{${"planId"}}`, encodeURIComponent(String(requestParameters.planId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete plan by UUID
     */
    async deletePlan(requestParameters: DeletePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deletePlanRaw(requestParameters, initOverrides);
    }

    /**
     * Update plan by UUID
     */
    async updatePlanRaw(requestParameters: UpdatePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasPlanPopulatedResponse>> {
        if (requestParameters.planId === null || requestParameters.planId === undefined) {
            throw new runtime.RequiredError('planId','Required parameter requestParameters.planId was null or undefined when calling updatePlan.');
        }

        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling updatePlan.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/plans/{planId}`.replace(`{${"planId"}}`, encodeURIComponent(String(requestParameters.planId))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasUpdatePlanPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasPlanPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Update plan by UUID
     */
    async updatePlan(requestParameters: UpdatePlanRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasPlanPopulatedResponse> {
        const response = await this.updatePlanRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
