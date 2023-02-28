import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class Category{
    @PrimaryGeneratedColumn()
    id: number;

    @Column({type: "citext"})
    category: string;

}