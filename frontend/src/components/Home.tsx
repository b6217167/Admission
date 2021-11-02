import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบรับผู้ป่วยใน</h1>
        <h4>Requirements</h4>
        <p>
          หากมีผู้ป่วยในที่ต้องการรักษาตัวอยู่ที่โรงพยาบาล เราจะใช้ระบบรับผู้ป่วยใน จัดหาห้องให้ผู้ป่วยในโดยหากมีโรคติดต่อจะบังคับให้อยู่ที่ห้อง 
         Negative pressure เท่านั้นและทำการบันทึกข้อมูลเข้าระบบ แต่ถ้าผู้ป่วยในไม่มีโรคติดต่อ 
         พยาบาลจะทำการสอบถามข้อมูลผู้ป่วยในว่าต้องการอยู่ห้องเดี่ยวหรือห้องรวม โดยหากต้องการอยู่ห้องเดี่ยว 
         จะมีให้เลือกอีก 2 ประเภท คือ ห้องเดี่ยวธรรมดา กับ ห้องเดี่ยวพิเศษ และเมื่อผู้ป่วยในเลือกห้องได้แล้ว 
         จะทำการเพิ่มข้อมูลของผู้ป่วยในเข้าไปในระบบรับผู้ป่วยใน 
        </p>
      </Container>
    </div>
  );
}
export default Home;