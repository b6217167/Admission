import { RoomInterface } from "./IRoom";
import { Right_TreatmentInterface } from "./IRight_Treatment";
import { PatientInterface } from "./IPatient";
import { RoomTypesInterface } from "./IRoomTypes";


export interface AdmissionInterface {
    ID : number,
    
	AdmitTime: Date,

	Right_TreatmentID: number,
	Right_Treatment: Right_TreatmentInterface,

	RoomID: number,
	Room: RoomInterface,

	RoomtypesID: number,
	Roomtypes: RoomTypesInterface,

	PatientID: number,
	Patient: PatientInterface
}