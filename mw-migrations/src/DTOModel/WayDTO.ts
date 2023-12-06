export type WayDTO = {
  uuid: string;
  name: string;
  dayReportUuids: string[];
  ownerUuid: string;
  monthReportUuids: string[];
  goalUuid: string;
  currentMentorUuids: string[];
  mentorRequestUuids: string[];
  isCompleted: boolean;
}