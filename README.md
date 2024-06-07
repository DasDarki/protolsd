# ProtoLSD
ProtoLSD is a custom DSL (domain-specific language) for creating protobuf definitions but simpler, more readable and more powerful.

## Configuration
ProtoLSDs behavior is configured by a configuration file. The configuration file should be placed at the root of the project and must be named `.protolsd.toml`.

## Specification
The ProtoLSD specification is based on the official protobuf specification version 3. Following are the additional features of ProtoLSD.

### File extension
Any file with the extension `.plsd` is considered a ProtoLSD file.

### Linebreaks, spaces and semicolons
ProtoLSD is not as strict as the official protobuf language when it comes to linebreaks, spaces and semicolons. ProtoLSD allows you to use linebreaks, spaces and semicolons wherever you want. ProtoLSD will automatically remove any unnecessary spaces and semicolons.

### Comments
ProtoLSD supports both single-line and multi-line comments. Single-line comments start with `//` and multi-line comments start with `/*` and end with `*/`.

### Preprocessor directives
ProtoLSD supports preprocessor directives. Preprocessor directives start with `#`. Preprocessor directives can have arguments which are passed in parentheses. If a preprocessor directive has no arguments, the parentheses can be omitted.

### Imports
ProtoLSD supports importing other ProtoLSD files and folders using the `import` keyword. The `import` keyword is followed by the path to the file to import. The path can be either an absolute path or a relative path. If the path is a relative path, it is relative to the directory of the importing file. Relative paths must start with `./` or `../`.   
ProtoLSD will resolve any imports recursively and generates the respective protobuf files and their includes. To disable this behavior, private imports can be used by using the `private` keyword before the `import` keyword, adding a `#private-import(global)` directive to the file, using the `#private-import` directive before the linebreak or semicolon, or using a bang at the end of the import path.

Valid import examples:
```proto
import "path/to/file.plsd";
private import "path/to/file.plsd";
import "path/to/file.plsd"!;
import "path/to/file.plsd" #private-import;
```
The `@` alias can be used to import files relative to the source folder.

#### Automatic imports
Every script which is in the same folder has access to all other scripts in the same folder. This means that you can import other scripts without specifying the path. This is useful for splitting up your proto file into multiple files. At compile time, ProtoLSD will automatically resolve the imports. Explicit imports will always take precedence over automatic imports. Importing folders makes the compiler aware of all files in that folder but only for that specific file. 

### Automatic field numbering
ProtoLSD automatically assigns field numbers to fields in their alphabetical order. You don't need to manually assign field numbers. If ProtoLSD "order_persist" mode is enabled, ProtoLSD tries to keep the field numbers the same across different versions of the proto file by storing needed information in a `.protolsd_persist` file in the same directory as the ProtoLSD configuration file.

Adding a bang `!` at the end of a number, will overwrite the current persisted field number.

### DataTypes
All protobuf data types are the same. There are some aliases for the data types. The following table shows the aliases for the data types:
| Alias | Data type |
|-------|-----------|
| `int` | `int32`   |
| `uint` | `uint32` |
| `sint` | `sint32` |
| `long` | `int64`  |
| `ulong` | `uint64`|

Additionally there are some small changes in the style of specific more advanced data types:
1. `repeated` fields can be defined by using the `[]` suffix after the data type. e.g. `int[] my_field = 1;`
2. `optional` fields can be defined by using the `?` suffix after the data type. e.g. `int? my_field = 1;`

Moreover there are data type aliases. By using the `#alias` directive, a data type can be aliased to another data type. This also works for messages:
```proto
#alias(int, int32)
#alias(MyMessage, MyOtherMessage)
```

### Enums
By default ProtoLSD changes the behaviour of enums. Enums are normally defined as a list of constants and will be generated like C++ enums. ProtoLSD defaults to C++ enum classes which are more context-safe. Given this ProtoLSD enum:
```proto	
enum MyEnum {
    FIRST = 0;
    SECOND = 1;
}
```

will generate this Protobuf 3 enum:
```proto
enum MyEnum {
    MyEnum_FIRST = 0;
    MyEnum_SECOND = 1;
}
```

To deactivate this behaviour, the `@NoEnumClass` annotation can be used prior to the enum definition:
```proto
@NoEnumClass
enum MyEnum {
    FIRST = 0;
    SECOND = 1;
}
```

Additionally, enums can be autoincremented by omitting any value:
```proto
enum MyEnum {
    FIRST;
    SECOND;
}
```
Here `FIRST` will be `0` and `SECOND` will be `1`.
Specifying a value for any of the constants will disable the autoincrementation for specific constant. The following constants will be autoincremented from the last specified value:
```proto
enum MyEnum {
    FIRST = 0;
    SECOND; // will be 1
    THIRD = 5;
    FOURTH; // will be 6
}
```

### Default request and response messages
Normally every protobuf service rpc must have a valid request and response message. ProtoLSD allows omitting request and response messages, if the rpc method does not require any. Therefore ProtoLSD will generate default request and response messages globally. These look like this:
```proto
message EmptyRequest {}
message EmptyResponse {}
```

Through this, ProtoLSD allows defining rpc methods without request and response messages:
```proto
service MyService {
    rpc MyMethod;
}
```

If a more expressive style is desired, empty request/response messages can be declared explicitly:
```proto
service MyService {
    rpc MyMethod(EmptyRequest) returns (EmptyResponse);
}
```

or slightly less verbose:
```proto
service MyService {
    rpc MyMethod() returns {};
}
```

Any combination of these styles is allowed.

### Multiple input parameters
ProtoLSD allows defining multiple input parameters for rpc methods, like this:
```proto
service MyService {
    rpc MyMethod(int param1, string param2);
}
```

Any rules for specifying message types apply to these parameters as well. Under the hood, these parameters will be transpiled to a single message type:
```proto
message MyMethodRequest {
    int param1 = 1;
    string param2 = 2;
}
```

Although it is better to explicitly define a name for the parameters, the parameter names are verbose and can be omitted:
```proto
service MyService {
    rpc MyMethod(int, string);
}
```

Additionally, parameters of the same type which are consecutive can be grouped together:
```proto
service MyService {
    rpc MyMethod(int, int, string);
}
```

can be simplified to:
```proto
service MyService {
    rpc MyMethod(int *2, string);
}
```

or when names are present, a plus sign can be used to group parameters:
```proto
service MyService {
    rpc MyMethod(int+ param1, param2, string);
}
```
Any named parameter after a grouped parameter will be assigned the same data type, until a new data type is specified.

The same works for return types:
```proto
service MyService {
    rpc MyMethod() returns (int, string);
}
```

### Importing external proto files
Its common that you need to import external protobuf files like google's protobuf libraries adding any or something like this. ProtoLSD doesn't directly support those imported types because it doesn't know about them. To solve this, ProtoLSD allows you declare such types through a respective preprocessor directive called `#decl-type`. This directive does not ensure the presence of the type, it just tells ProtoLSD that the type exists, which disables any type checking for this type. The directive can be used like this:
```proto
#decl-type(google.protobuf.Any)
```

To ensure, that the type is present after generating the protobuf files, an import statement should be added to the file, importing the protobuf file. This can be done by using the `import` statement:
```proto
import "google/protobuf/any.proto";
```

### Protbuf 3 compatibility
Currently, not all functionality of Protobuf 3 is supported by ProtoLSD. The following features are not supported:
- Oneof
- Extensions
- Services with streaming
- Possibly more

Currently there is also limited support of nested packages/messages. For nested messages there is no issue but for nested packages, only the depth of 1 is supprted:
```proto
somepackage.Message // this is supported
somepackage.subpackage.Message // this is not supported
```

### Some notes to the persist order mode
This only works for directly defined messages. This does not work for messages which are defined through direct input/return in the method definition.