import {IconProps} from "src/component/icon/Icon";

/**
 * Upload icon
 */
export const UploadIcon = (props: IconProps) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="none"
      className={props.className}
      strokeWidth="1"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
      <polyline points="17 8 12 3 7 8" />
      <line
        x1="12"
        y1="3"
        x2="12"
        y2="15"
      />
    </svg>
  );
};