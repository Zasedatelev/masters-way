import path from "path";
import * as fs from "fs";

const OUT_DIR = "migration-logs";

const date = new Date().toISOString();
const logsDir = path.join(process.cwd(), OUT_DIR);

export const logToFile = (textToLog: string, fileName: string) => {
  if (!fs.existsSync(logsDir)) {
    fs.mkdirSync(logsDir);
  }

  const logFilePath = path.join(logsDir, `${fileName}_${date}.txt`);
  
  if (fs.existsSync(logFilePath)) {
    console.log(textToLog);
    fs.appendFileSync(logFilePath, `${textToLog}\n`, "utf8");
  } else {
    console.log('file created')
    console.log(textToLog);
    fs.writeFileSync(logFilePath, textToLog);
    fs.appendFileSync(logFilePath, `$\n`, "utf8");
  }
}
