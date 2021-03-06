/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
package com.worldpay.innovation.wpwithin.rpc.types;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SuppressWarnings({"cast", "rawtypes", "serial", "unchecked"})
@Generated(value = "Autogenerated by Thrift Compiler (0.9.3)", date = "2016-09-08")
public class Price implements org.apache.thrift.TBase<Price, Price._Fields>, java.io.Serializable, Cloneable, Comparable<Price> {
  private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("Price");

  private static final org.apache.thrift.protocol.TField ID_FIELD_DESC = new org.apache.thrift.protocol.TField("id", org.apache.thrift.protocol.TType.I32, (short)1);
  private static final org.apache.thrift.protocol.TField DESCRIPTION_FIELD_DESC = new org.apache.thrift.protocol.TField("description", org.apache.thrift.protocol.TType.STRING, (short)2);
  private static final org.apache.thrift.protocol.TField PRICE_PER_UNIT_FIELD_DESC = new org.apache.thrift.protocol.TField("pricePerUnit", org.apache.thrift.protocol.TType.STRUCT, (short)3);
  private static final org.apache.thrift.protocol.TField UNIT_ID_FIELD_DESC = new org.apache.thrift.protocol.TField("unitId", org.apache.thrift.protocol.TType.I32, (short)4);
  private static final org.apache.thrift.protocol.TField UNIT_DESCRIPTION_FIELD_DESC = new org.apache.thrift.protocol.TField("unitDescription", org.apache.thrift.protocol.TType.STRING, (short)5);

  private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
  static {
    schemes.put(StandardScheme.class, new PriceStandardSchemeFactory());
    schemes.put(TupleScheme.class, new PriceTupleSchemeFactory());
  }

  public int id; // required
  public String description; // required
  public PricePerUnit pricePerUnit; // required
  public int unitId; // required
  public String unitDescription; // required

  /** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
  public enum _Fields implements org.apache.thrift.TFieldIdEnum {
    ID((short)1, "id"),
    DESCRIPTION((short)2, "description"),
    PRICE_PER_UNIT((short)3, "pricePerUnit"),
    UNIT_ID((short)4, "unitId"),
    UNIT_DESCRIPTION((short)5, "unitDescription");

    private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

    static {
      for (_Fields field : EnumSet.allOf(_Fields.class)) {
        byName.put(field.getFieldName(), field);
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, or null if its not found.
     */
    public static _Fields findByThriftId(int fieldId) {
      switch(fieldId) {
        case 1: // ID
          return ID;
        case 2: // DESCRIPTION
          return DESCRIPTION;
        case 3: // PRICE_PER_UNIT
          return PRICE_PER_UNIT;
        case 4: // UNIT_ID
          return UNIT_ID;
        case 5: // UNIT_DESCRIPTION
          return UNIT_DESCRIPTION;
        default:
          return null;
      }
    }

    /**
     * Find the _Fields constant that matches fieldId, throwing an exception
     * if it is not found.
     */
    public static _Fields findByThriftIdOrThrow(int fieldId) {
      _Fields fields = findByThriftId(fieldId);
      if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
      return fields;
    }

    /**
     * Find the _Fields constant that matches name, or null if its not found.
     */
    public static _Fields findByName(String name) {
      return byName.get(name);
    }

    private final short _thriftId;
    private final String _fieldName;

    _Fields(short thriftId, String fieldName) {
      _thriftId = thriftId;
      _fieldName = fieldName;
    }

    public short getThriftFieldId() {
      return _thriftId;
    }

    public String getFieldName() {
      return _fieldName;
    }
  }

  // isset id assignments
  private static final int __ID_ISSET_ID = 0;
  private static final int __UNITID_ISSET_ID = 1;
  private byte __isset_bitfield = 0;
  public static final Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> metaDataMap;
  static {
    Map<_Fields, org.apache.thrift.meta_data.FieldMetaData> tmpMap = new EnumMap<_Fields, org.apache.thrift.meta_data.FieldMetaData>(_Fields.class);
    tmpMap.put(_Fields.ID, new org.apache.thrift.meta_data.FieldMetaData("id", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32)));
    tmpMap.put(_Fields.DESCRIPTION, new org.apache.thrift.meta_data.FieldMetaData("description", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    tmpMap.put(_Fields.PRICE_PER_UNIT, new org.apache.thrift.meta_data.FieldMetaData("pricePerUnit", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRUCT        , "PricePerUnit")));
    tmpMap.put(_Fields.UNIT_ID, new org.apache.thrift.meta_data.FieldMetaData("unitId", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.I32)));
    tmpMap.put(_Fields.UNIT_DESCRIPTION, new org.apache.thrift.meta_data.FieldMetaData("unitDescription", org.apache.thrift.TFieldRequirementType.DEFAULT, 
        new org.apache.thrift.meta_data.FieldValueMetaData(org.apache.thrift.protocol.TType.STRING)));
    metaDataMap = Collections.unmodifiableMap(tmpMap);
    org.apache.thrift.meta_data.FieldMetaData.addStructMetaDataMap(Price.class, metaDataMap);
  }

  public Price() {
  }

  public Price(
    int id,
    String description,
    PricePerUnit pricePerUnit,
    int unitId,
    String unitDescription)
  {
    this();
    this.id = id;
    setIdIsSet(true);
    this.description = description;
    this.pricePerUnit = pricePerUnit;
    this.unitId = unitId;
    setUnitIdIsSet(true);
    this.unitDescription = unitDescription;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public Price(Price other) {
    __isset_bitfield = other.__isset_bitfield;
    this.id = other.id;
    if (other.isSetDescription()) {
      this.description = other.description;
    }
    if (other.isSetPricePerUnit()) {
      this.pricePerUnit = other.pricePerUnit;
    }
    this.unitId = other.unitId;
    if (other.isSetUnitDescription()) {
      this.unitDescription = other.unitDescription;
    }
  }

  public Price deepCopy() {
    return new Price(this);
  }

  @Override
  public void clear() {
    setIdIsSet(false);
    this.id = 0;
    this.description = null;
    this.pricePerUnit = null;
    setUnitIdIsSet(false);
    this.unitId = 0;
    this.unitDescription = null;
  }

  public int getId() {
    return this.id;
  }

  public Price setId(int id) {
    this.id = id;
    setIdIsSet(true);
    return this;
  }

  public void unsetId() {
    __isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __ID_ISSET_ID);
  }

  /** Returns true if field id is set (has been assigned a value) and false otherwise */
  public boolean isSetId() {
    return EncodingUtils.testBit(__isset_bitfield, __ID_ISSET_ID);
  }

  public void setIdIsSet(boolean value) {
    __isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __ID_ISSET_ID, value);
  }

  public String getDescription() {
    return this.description;
  }

  public Price setDescription(String description) {
    this.description = description;
    return this;
  }

  public void unsetDescription() {
    this.description = null;
  }

  /** Returns true if field description is set (has been assigned a value) and false otherwise */
  public boolean isSetDescription() {
    return this.description != null;
  }

  public void setDescriptionIsSet(boolean value) {
    if (!value) {
      this.description = null;
    }
  }

  public PricePerUnit getPricePerUnit() {
    return this.pricePerUnit;
  }

  public Price setPricePerUnit(PricePerUnit pricePerUnit) {
    this.pricePerUnit = pricePerUnit;
    return this;
  }

  public void unsetPricePerUnit() {
    this.pricePerUnit = null;
  }

  /** Returns true if field pricePerUnit is set (has been assigned a value) and false otherwise */
  public boolean isSetPricePerUnit() {
    return this.pricePerUnit != null;
  }

  public void setPricePerUnitIsSet(boolean value) {
    if (!value) {
      this.pricePerUnit = null;
    }
  }

  public int getUnitId() {
    return this.unitId;
  }

  public Price setUnitId(int unitId) {
    this.unitId = unitId;
    setUnitIdIsSet(true);
    return this;
  }

  public void unsetUnitId() {
    __isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __UNITID_ISSET_ID);
  }

  /** Returns true if field unitId is set (has been assigned a value) and false otherwise */
  public boolean isSetUnitId() {
    return EncodingUtils.testBit(__isset_bitfield, __UNITID_ISSET_ID);
  }

  public void setUnitIdIsSet(boolean value) {
    __isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __UNITID_ISSET_ID, value);
  }

  public String getUnitDescription() {
    return this.unitDescription;
  }

  public Price setUnitDescription(String unitDescription) {
    this.unitDescription = unitDescription;
    return this;
  }

  public void unsetUnitDescription() {
    this.unitDescription = null;
  }

  /** Returns true if field unitDescription is set (has been assigned a value) and false otherwise */
  public boolean isSetUnitDescription() {
    return this.unitDescription != null;
  }

  public void setUnitDescriptionIsSet(boolean value) {
    if (!value) {
      this.unitDescription = null;
    }
  }

  public void setFieldValue(_Fields field, Object value) {
    switch (field) {
    case ID:
      if (value == null) {
        unsetId();
      } else {
        setId((Integer)value);
      }
      break;

    case DESCRIPTION:
      if (value == null) {
        unsetDescription();
      } else {
        setDescription((String)value);
      }
      break;

    case PRICE_PER_UNIT:
      if (value == null) {
        unsetPricePerUnit();
      } else {
        setPricePerUnit((PricePerUnit)value);
      }
      break;

    case UNIT_ID:
      if (value == null) {
        unsetUnitId();
      } else {
        setUnitId((Integer)value);
      }
      break;

    case UNIT_DESCRIPTION:
      if (value == null) {
        unsetUnitDescription();
      } else {
        setUnitDescription((String)value);
      }
      break;

    }
  }

  public Object getFieldValue(_Fields field) {
    switch (field) {
    case ID:
      return getId();

    case DESCRIPTION:
      return getDescription();

    case PRICE_PER_UNIT:
      return getPricePerUnit();

    case UNIT_ID:
      return getUnitId();

    case UNIT_DESCRIPTION:
      return getUnitDescription();

    }
    throw new IllegalStateException();
  }

  /** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
  public boolean isSet(_Fields field) {
    if (field == null) {
      throw new IllegalArgumentException();
    }

    switch (field) {
    case ID:
      return isSetId();
    case DESCRIPTION:
      return isSetDescription();
    case PRICE_PER_UNIT:
      return isSetPricePerUnit();
    case UNIT_ID:
      return isSetUnitId();
    case UNIT_DESCRIPTION:
      return isSetUnitDescription();
    }
    throw new IllegalStateException();
  }

  @Override
  public boolean equals(Object that) {
    if (that == null)
      return false;
    if (that instanceof Price)
      return this.equals((Price)that);
    return false;
  }

  public boolean equals(Price that) {
    if (that == null)
      return false;

    boolean this_present_id = true;
    boolean that_present_id = true;
    if (this_present_id || that_present_id) {
      if (!(this_present_id && that_present_id))
        return false;
      if (this.id != that.id)
        return false;
    }

    boolean this_present_description = true && this.isSetDescription();
    boolean that_present_description = true && that.isSetDescription();
    if (this_present_description || that_present_description) {
      if (!(this_present_description && that_present_description))
        return false;
      if (!this.description.equals(that.description))
        return false;
    }

    boolean this_present_pricePerUnit = true && this.isSetPricePerUnit();
    boolean that_present_pricePerUnit = true && that.isSetPricePerUnit();
    if (this_present_pricePerUnit || that_present_pricePerUnit) {
      if (!(this_present_pricePerUnit && that_present_pricePerUnit))
        return false;
      if (!this.pricePerUnit.equals(that.pricePerUnit))
        return false;
    }

    boolean this_present_unitId = true;
    boolean that_present_unitId = true;
    if (this_present_unitId || that_present_unitId) {
      if (!(this_present_unitId && that_present_unitId))
        return false;
      if (this.unitId != that.unitId)
        return false;
    }

    boolean this_present_unitDescription = true && this.isSetUnitDescription();
    boolean that_present_unitDescription = true && that.isSetUnitDescription();
    if (this_present_unitDescription || that_present_unitDescription) {
      if (!(this_present_unitDescription && that_present_unitDescription))
        return false;
      if (!this.unitDescription.equals(that.unitDescription))
        return false;
    }

    return true;
  }

  @Override
  public int hashCode() {
    List<Object> list = new ArrayList<Object>();

    boolean present_id = true;
    list.add(present_id);
    if (present_id)
      list.add(id);

    boolean present_description = true && (isSetDescription());
    list.add(present_description);
    if (present_description)
      list.add(description);

    boolean present_pricePerUnit = true && (isSetPricePerUnit());
    list.add(present_pricePerUnit);
    if (present_pricePerUnit)
      list.add(pricePerUnit);

    boolean present_unitId = true;
    list.add(present_unitId);
    if (present_unitId)
      list.add(unitId);

    boolean present_unitDescription = true && (isSetUnitDescription());
    list.add(present_unitDescription);
    if (present_unitDescription)
      list.add(unitDescription);

    return list.hashCode();
  }

  @Override
  public int compareTo(Price other) {
    if (!getClass().equals(other.getClass())) {
      return getClass().getName().compareTo(other.getClass().getName());
    }

    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetId()).compareTo(other.isSetId());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetId()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.id, other.id);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetDescription()).compareTo(other.isSetDescription());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetDescription()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.description, other.description);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetPricePerUnit()).compareTo(other.isSetPricePerUnit());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetPricePerUnit()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.pricePerUnit, other.pricePerUnit);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetUnitId()).compareTo(other.isSetUnitId());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetUnitId()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.unitId, other.unitId);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    lastComparison = Boolean.valueOf(isSetUnitDescription()).compareTo(other.isSetUnitDescription());
    if (lastComparison != 0) {
      return lastComparison;
    }
    if (isSetUnitDescription()) {
      lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.unitDescription, other.unitDescription);
      if (lastComparison != 0) {
        return lastComparison;
      }
    }
    return 0;
  }

  public _Fields fieldForId(int fieldId) {
    return _Fields.findByThriftId(fieldId);
  }

  public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
    schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
  }

  public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
    schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder("Price(");
    boolean first = true;

    sb.append("id:");
    sb.append(this.id);
    first = false;
    if (!first) sb.append(", ");
    sb.append("description:");
    if (this.description == null) {
      sb.append("null");
    } else {
      sb.append(this.description);
    }
    first = false;
    if (!first) sb.append(", ");
    sb.append("pricePerUnit:");
    if (this.pricePerUnit == null) {
      sb.append("null");
    } else {
      sb.append(this.pricePerUnit);
    }
    first = false;
    if (!first) sb.append(", ");
    sb.append("unitId:");
    sb.append(this.unitId);
    first = false;
    if (!first) sb.append(", ");
    sb.append("unitDescription:");
    if (this.unitDescription == null) {
      sb.append("null");
    } else {
      sb.append(this.unitDescription);
    }
    first = false;
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws org.apache.thrift.TException {
    // check for required fields
    // check for sub-struct validity
  }

  private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
    try {
      write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
    try {
      // it doesn't seem like you should have to do this, but java serialization is wacky, and doesn't call the default constructor.
      __isset_bitfield = 0;
      read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
    } catch (org.apache.thrift.TException te) {
      throw new java.io.IOException(te);
    }
  }

  private static class PriceStandardSchemeFactory implements SchemeFactory {
    public PriceStandardScheme getScheme() {
      return new PriceStandardScheme();
    }
  }

  private static class PriceStandardScheme extends StandardScheme<Price> {

    public void read(org.apache.thrift.protocol.TProtocol iprot, Price struct) throws org.apache.thrift.TException {
      org.apache.thrift.protocol.TField schemeField;
      iprot.readStructBegin();
      while (true)
      {
        schemeField = iprot.readFieldBegin();
        if (schemeField.type == org.apache.thrift.protocol.TType.STOP) { 
          break;
        }
        switch (schemeField.id) {
          case 1: // ID
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.id = iprot.readI32();
              struct.setIdIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 2: // DESCRIPTION
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.description = iprot.readString();
              struct.setDescriptionIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 3: // PRICE_PER_UNIT
            if (schemeField.type == org.apache.thrift.protocol.TType.STRUCT) {
              struct.pricePerUnit = new PricePerUnit();
              struct.pricePerUnit.read(iprot);
              struct.setPricePerUnitIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 4: // UNIT_ID
            if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
              struct.unitId = iprot.readI32();
              struct.setUnitIdIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          case 5: // UNIT_DESCRIPTION
            if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
              struct.unitDescription = iprot.readString();
              struct.setUnitDescriptionIsSet(true);
            } else { 
              org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
            }
            break;
          default:
            org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
        }
        iprot.readFieldEnd();
      }
      iprot.readStructEnd();

      // check for required fields of primitive type, which can't be checked in the validate method
      struct.validate();
    }

    public void write(org.apache.thrift.protocol.TProtocol oprot, Price struct) throws org.apache.thrift.TException {
      struct.validate();

      oprot.writeStructBegin(STRUCT_DESC);
      oprot.writeFieldBegin(ID_FIELD_DESC);
      oprot.writeI32(struct.id);
      oprot.writeFieldEnd();
      if (struct.description != null) {
        oprot.writeFieldBegin(DESCRIPTION_FIELD_DESC);
        oprot.writeString(struct.description);
        oprot.writeFieldEnd();
      }
      if (struct.pricePerUnit != null) {
        oprot.writeFieldBegin(PRICE_PER_UNIT_FIELD_DESC);
        struct.pricePerUnit.write(oprot);
        oprot.writeFieldEnd();
      }
      oprot.writeFieldBegin(UNIT_ID_FIELD_DESC);
      oprot.writeI32(struct.unitId);
      oprot.writeFieldEnd();
      if (struct.unitDescription != null) {
        oprot.writeFieldBegin(UNIT_DESCRIPTION_FIELD_DESC);
        oprot.writeString(struct.unitDescription);
        oprot.writeFieldEnd();
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

  }

  private static class PriceTupleSchemeFactory implements SchemeFactory {
    public PriceTupleScheme getScheme() {
      return new PriceTupleScheme();
    }
  }

  private static class PriceTupleScheme extends TupleScheme<Price> {

    @Override
    public void write(org.apache.thrift.protocol.TProtocol prot, Price struct) throws org.apache.thrift.TException {
      TTupleProtocol oprot = (TTupleProtocol) prot;
      BitSet optionals = new BitSet();
      if (struct.isSetId()) {
        optionals.set(0);
      }
      if (struct.isSetDescription()) {
        optionals.set(1);
      }
      if (struct.isSetPricePerUnit()) {
        optionals.set(2);
      }
      if (struct.isSetUnitId()) {
        optionals.set(3);
      }
      if (struct.isSetUnitDescription()) {
        optionals.set(4);
      }
      oprot.writeBitSet(optionals, 5);
      if (struct.isSetId()) {
        oprot.writeI32(struct.id);
      }
      if (struct.isSetDescription()) {
        oprot.writeString(struct.description);
      }
      if (struct.isSetPricePerUnit()) {
        struct.pricePerUnit.write(oprot);
      }
      if (struct.isSetUnitId()) {
        oprot.writeI32(struct.unitId);
      }
      if (struct.isSetUnitDescription()) {
        oprot.writeString(struct.unitDescription);
      }
    }

    @Override
    public void read(org.apache.thrift.protocol.TProtocol prot, Price struct) throws org.apache.thrift.TException {
      TTupleProtocol iprot = (TTupleProtocol) prot;
      BitSet incoming = iprot.readBitSet(5);
      if (incoming.get(0)) {
        struct.id = iprot.readI32();
        struct.setIdIsSet(true);
      }
      if (incoming.get(1)) {
        struct.description = iprot.readString();
        struct.setDescriptionIsSet(true);
      }
      if (incoming.get(2)) {
        struct.pricePerUnit = new PricePerUnit();
        struct.pricePerUnit.read(iprot);
        struct.setPricePerUnitIsSet(true);
      }
      if (incoming.get(3)) {
        struct.unitId = iprot.readI32();
        struct.setUnitIdIsSet(true);
      }
      if (incoming.get(4)) {
        struct.unitDescription = iprot.readString();
        struct.setUnitDescriptionIsSet(true);
      }
    }
  }

}

