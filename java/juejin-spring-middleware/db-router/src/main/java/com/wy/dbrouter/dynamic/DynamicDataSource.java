package com.wy.dbrouter.dynamic;

import com.wy.dbrouter.DBContextHolder;
import org.springframework.jdbc.datasource.lookup.AbstractRoutingDataSource;


public class DynamicDataSource extends AbstractRoutingDataSource {

    @Override
    protected Object determineCurrentLookupKey() {
        return "db"+ DBContextHolder.getDBKey();
    }
}
