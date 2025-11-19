import { PrismaLibSql } from "@prisma/adapter-libsql";
import { PrismaClient } from "./prisma/client";

const adapter = new PrismaLibSql({
	url: `${process.env.DATABASE_URL}`,
	authToken: `${process.env.DATABASE_TOKEN}`,
});

export default new PrismaClient({ adapter });
