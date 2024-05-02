import {
  CreateMetricRequest,
  DeleteMetricRequest,
  SchemasMetricResponse,
  UpdateMetricRequest,
} from "src/apiAutogenerated";
import {metricService} from "src/service/services";

/**
 * Provides methods to interact with the metrics
 */
export class MetricService {

  /**
   * Create metric
   */
  public static async createMetric(requestParameters: CreateMetricRequest): Promise<SchemasMetricResponse> {
    const jobDone = await metricService.createMetric(requestParameters);

    return jobDone;
  }

  /**
   * Update metric
   */
  public static async updateMetric(requestParameters: UpdateMetricRequest): Promise<SchemasMetricResponse> {
    const updatedJobDone = await metricService.updateMetric(requestParameters);

    return updatedJobDone;
  }

  /**
   * Delete metric by UUID
   */
  public static async deleteMetric(requestParameters: DeleteMetricRequest): Promise<void> {
    await metricService.deleteMetric(requestParameters);
  }

}