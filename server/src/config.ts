import { TypeOrmModuleOptions } from '@nestjs/typeorm';

interface Config {
  db: TypeOrmModuleOptions;
}

export default (): Config => ({
  db: {
    database: process.env.POSTGRES_DATABASE,
    username: process.env.POSTGRES_USER,
    password: process.env.POSTGRES_PASSWORD,
    host: process.env.POSTGRES_HOST,
    type: 'postgres',
    synchronize: process.env.NODE_ENV === 'production',
  },
});
