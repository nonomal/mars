import * as $protobuf from "protobufjs";
/** Namespace auth. */
export namespace auth {

    /** Properties of a LoginRequest. */
    interface ILoginRequest {

        /** LoginRequest username */
        username?: (string|null);

        /** LoginRequest password */
        password?: (string|null);
    }

    /** Represents a LoginRequest. */
    class LoginRequest implements ILoginRequest {

        /**
         * Constructs a new LoginRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.ILoginRequest);

        /** LoginRequest username. */
        public username: string;

        /** LoginRequest password. */
        public password: string;

        /**
         * Encodes the specified LoginRequest message. Does not implicitly {@link auth.LoginRequest.verify|verify} messages.
         * @param message LoginRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.LoginRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a LoginRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns LoginRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.LoginRequest;
    }

    /** Properties of a LoginResponse. */
    interface ILoginResponse {

        /** LoginResponse token */
        token?: (string|null);

        /** LoginResponse expires_in */
        expires_in?: (number|null);
    }

    /** Represents a LoginResponse. */
    class LoginResponse implements ILoginResponse {

        /**
         * Constructs a new LoginResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.ILoginResponse);

        /** LoginResponse token. */
        public token: string;

        /** LoginResponse expires_in. */
        public expires_in: number;

        /**
         * Encodes the specified LoginResponse message. Does not implicitly {@link auth.LoginResponse.verify|verify} messages.
         * @param message LoginResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.LoginResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a LoginResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns LoginResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.LoginResponse;
    }

    /** Properties of an ExchangeRequest. */
    interface IExchangeRequest {

        /** ExchangeRequest code */
        code?: (string|null);
    }

    /** Represents an ExchangeRequest. */
    class ExchangeRequest implements IExchangeRequest {

        /**
         * Constructs a new ExchangeRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.IExchangeRequest);

        /** ExchangeRequest code. */
        public code: string;

        /**
         * Encodes the specified ExchangeRequest message. Does not implicitly {@link auth.ExchangeRequest.verify|verify} messages.
         * @param message ExchangeRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.ExchangeRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExchangeRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExchangeRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.ExchangeRequest;
    }

    /** Properties of an ExchangeResponse. */
    interface IExchangeResponse {

        /** ExchangeResponse token */
        token?: (string|null);

        /** ExchangeResponse expires_in */
        expires_in?: (number|null);
    }

    /** Represents an ExchangeResponse. */
    class ExchangeResponse implements IExchangeResponse {

        /**
         * Constructs a new ExchangeResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.IExchangeResponse);

        /** ExchangeResponse token. */
        public token: string;

        /** ExchangeResponse expires_in. */
        public expires_in: number;

        /**
         * Encodes the specified ExchangeResponse message. Does not implicitly {@link auth.ExchangeResponse.verify|verify} messages.
         * @param message ExchangeResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.ExchangeResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExchangeResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExchangeResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.ExchangeResponse;
    }

    /** Properties of an InfoRequest. */
    interface IInfoRequest {
    }

    /** Represents an InfoRequest. */
    class InfoRequest implements IInfoRequest {

        /**
         * Constructs a new InfoRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.IInfoRequest);

        /**
         * Encodes the specified InfoRequest message. Does not implicitly {@link auth.InfoRequest.verify|verify} messages.
         * @param message InfoRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.InfoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InfoRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InfoRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.InfoRequest;
    }

    /** Properties of an InfoResponse. */
    interface IInfoResponse {

        /** InfoResponse id */
        id?: (string|null);

        /** InfoResponse avatar */
        avatar?: (string|null);

        /** InfoResponse name */
        name?: (string|null);

        /** InfoResponse email */
        email?: (string|null);

        /** InfoResponse logout_url */
        logout_url?: (string|null);

        /** InfoResponse roles */
        roles?: (string[]|null);
    }

    /** Represents an InfoResponse. */
    class InfoResponse implements IInfoResponse {

        /**
         * Constructs a new InfoResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.IInfoResponse);

        /** InfoResponse id. */
        public id: string;

        /** InfoResponse avatar. */
        public avatar: string;

        /** InfoResponse name. */
        public name: string;

        /** InfoResponse email. */
        public email: string;

        /** InfoResponse logout_url. */
        public logout_url: string;

        /** InfoResponse roles. */
        public roles: string[];

        /**
         * Encodes the specified InfoResponse message. Does not implicitly {@link auth.InfoResponse.verify|verify} messages.
         * @param message InfoResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.InfoResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InfoResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InfoResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.InfoResponse;
    }

    /** Properties of a SettingsRequest. */
    interface ISettingsRequest {
    }

    /** Represents a SettingsRequest. */
    class SettingsRequest implements ISettingsRequest {

        /**
         * Constructs a new SettingsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.ISettingsRequest);

        /**
         * Encodes the specified SettingsRequest message. Does not implicitly {@link auth.SettingsRequest.verify|verify} messages.
         * @param message SettingsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.SettingsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a SettingsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns SettingsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.SettingsRequest;
    }

    /** Properties of a SettingsResponse. */
    interface ISettingsResponse {

        /** SettingsResponse items */
        items?: (auth.SettingsResponse.OidcSetting[]|null);
    }

    /** Represents a SettingsResponse. */
    class SettingsResponse implements ISettingsResponse {

        /**
         * Constructs a new SettingsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: auth.ISettingsResponse);

        /** SettingsResponse items. */
        public items: auth.SettingsResponse.OidcSetting[];

        /**
         * Encodes the specified SettingsResponse message. Does not implicitly {@link auth.SettingsResponse.verify|verify} messages.
         * @param message SettingsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: auth.SettingsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a SettingsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns SettingsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.SettingsResponse;
    }

    namespace SettingsResponse {

        /** Properties of an OidcSetting. */
        interface IOidcSetting {

            /** OidcSetting enabled */
            enabled?: (boolean|null);

            /** OidcSetting name */
            name?: (string|null);

            /** OidcSetting url */
            url?: (string|null);

            /** OidcSetting end_session_endpoint */
            end_session_endpoint?: (string|null);

            /** OidcSetting state */
            state?: (string|null);
        }

        /** Represents an OidcSetting. */
        class OidcSetting implements IOidcSetting {

            /**
             * Constructs a new OidcSetting.
             * @param [properties] Properties to set
             */
            constructor(properties?: auth.SettingsResponse.IOidcSetting);

            /** OidcSetting enabled. */
            public enabled: boolean;

            /** OidcSetting name. */
            public name: string;

            /** OidcSetting url. */
            public url: string;

            /** OidcSetting end_session_endpoint. */
            public end_session_endpoint: string;

            /** OidcSetting state. */
            public state: string;

            /**
             * Encodes the specified OidcSetting message. Does not implicitly {@link auth.SettingsResponse.OidcSetting.verify|verify} messages.
             * @param message OidcSetting message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: auth.SettingsResponse.OidcSetting, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an OidcSetting message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns OidcSetting
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.SettingsResponse.OidcSetting;
        }
    }

    /** Represents an Auth */
    class Auth extends $protobuf.rpc.Service {

        /**
         * Constructs a new Auth service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls Login.
         * @param request LoginRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and LoginResponse
         */
        public login(request: auth.LoginRequest, callback: auth.Auth.LoginCallback): void;

        /**
         * Calls Login.
         * @param request LoginRequest message or plain object
         * @returns Promise
         */
        public login(request: auth.LoginRequest): Promise<auth.LoginResponse>;

        /**
         * Calls Info.
         * @param request InfoRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and InfoResponse
         */
        public info(request: auth.InfoRequest, callback: auth.Auth.InfoCallback): void;

        /**
         * Calls Info.
         * @param request InfoRequest message or plain object
         * @returns Promise
         */
        public info(request: auth.InfoRequest): Promise<auth.InfoResponse>;

        /**
         * Calls Settings.
         * @param request SettingsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and SettingsResponse
         */
        public settings(request: auth.SettingsRequest, callback: auth.Auth.SettingsCallback): void;

        /**
         * Calls Settings.
         * @param request SettingsRequest message or plain object
         * @returns Promise
         */
        public settings(request: auth.SettingsRequest): Promise<auth.SettingsResponse>;

        /**
         * Calls Exchange.
         * @param request ExchangeRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ExchangeResponse
         */
        public exchange(request: auth.ExchangeRequest, callback: auth.Auth.ExchangeCallback): void;

        /**
         * Calls Exchange.
         * @param request ExchangeRequest message or plain object
         * @returns Promise
         */
        public exchange(request: auth.ExchangeRequest): Promise<auth.ExchangeResponse>;
    }

    namespace Auth {

        /**
         * Callback as used by {@link auth.Auth#login}.
         * @param error Error, if any
         * @param [response] LoginResponse
         */
        type LoginCallback = (error: (Error|null), response?: auth.LoginResponse) => void;

        /**
         * Callback as used by {@link auth.Auth#info}.
         * @param error Error, if any
         * @param [response] InfoResponse
         */
        type InfoCallback = (error: (Error|null), response?: auth.InfoResponse) => void;

        /**
         * Callback as used by {@link auth.Auth#settings}.
         * @param error Error, if any
         * @param [response] SettingsResponse
         */
        type SettingsCallback = (error: (Error|null), response?: auth.SettingsResponse) => void;

        /**
         * Callback as used by {@link auth.Auth#exchange}.
         * @param error Error, if any
         * @param [response] ExchangeResponse
         */
        type ExchangeCallback = (error: (Error|null), response?: auth.ExchangeResponse) => void;
    }
}

/** Namespace google. */
export namespace google {

    /** Namespace api. */
    namespace api {

        /** Properties of a Http. */
        interface IHttp {

            /** Http rules */
            rules?: (google.api.HttpRule[]|null);
        }

        /** Represents a Http. */
        class Http implements IHttp {

            /**
             * Constructs a new Http.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.api.IHttp);

            /** Http rules. */
            public rules: google.api.HttpRule[];

            /**
             * Encodes the specified Http message. Does not implicitly {@link google.api.Http.verify|verify} messages.
             * @param message Http message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.api.Http, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a Http message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Http
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.api.Http;
        }

        /** Properties of a HttpRule. */
        interface IHttpRule {

            /** HttpRule get */
            get?: (string|null);

            /** HttpRule put */
            put?: (string|null);

            /** HttpRule post */
            post?: (string|null);

            /** HttpRule delete */
            "delete"?: (string|null);

            /** HttpRule patch */
            patch?: (string|null);

            /** HttpRule custom */
            custom?: (google.api.CustomHttpPattern|null);

            /** HttpRule selector */
            selector?: (string|null);

            /** HttpRule body */
            body?: (string|null);

            /** HttpRule additional_bindings */
            additional_bindings?: (google.api.HttpRule[]|null);
        }

        /** Represents a HttpRule. */
        class HttpRule implements IHttpRule {

            /**
             * Constructs a new HttpRule.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.api.IHttpRule);

            /** HttpRule get. */
            public get?: (string|null);

            /** HttpRule put. */
            public put?: (string|null);

            /** HttpRule post. */
            public post?: (string|null);

            /** HttpRule delete. */
            public delete?: (string|null);

            /** HttpRule patch. */
            public patch?: (string|null);

            /** HttpRule custom. */
            public custom?: (google.api.CustomHttpPattern|null);

            /** HttpRule selector. */
            public selector: string;

            /** HttpRule body. */
            public body: string;

            /** HttpRule additional_bindings. */
            public additional_bindings: google.api.HttpRule[];

            /** HttpRule pattern. */
            public pattern?: ("get"|"put"|"post"|"delete"|"patch"|"custom");

            /**
             * Encodes the specified HttpRule message. Does not implicitly {@link google.api.HttpRule.verify|verify} messages.
             * @param message HttpRule message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.api.HttpRule, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a HttpRule message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns HttpRule
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.api.HttpRule;
        }

        /** Properties of a CustomHttpPattern. */
        interface ICustomHttpPattern {

            /** CustomHttpPattern kind */
            kind?: (string|null);

            /** CustomHttpPattern path */
            path?: (string|null);
        }

        /** Represents a CustomHttpPattern. */
        class CustomHttpPattern implements ICustomHttpPattern {

            /**
             * Constructs a new CustomHttpPattern.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.api.ICustomHttpPattern);

            /** CustomHttpPattern kind. */
            public kind: string;

            /** CustomHttpPattern path. */
            public path: string;

            /**
             * Encodes the specified CustomHttpPattern message. Does not implicitly {@link google.api.CustomHttpPattern.verify|verify} messages.
             * @param message CustomHttpPattern message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.api.CustomHttpPattern, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a CustomHttpPattern message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns CustomHttpPattern
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.api.CustomHttpPattern;
        }
    }

    /** Namespace protobuf. */
    namespace protobuf {

        /** Properties of a FileDescriptorSet. */
        interface IFileDescriptorSet {

            /** FileDescriptorSet file */
            file?: (google.protobuf.FileDescriptorProto[]|null);
        }

        /** Represents a FileDescriptorSet. */
        class FileDescriptorSet implements IFileDescriptorSet {

            /**
             * Constructs a new FileDescriptorSet.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IFileDescriptorSet);

            /** FileDescriptorSet file. */
            public file: google.protobuf.FileDescriptorProto[];

            /**
             * Encodes the specified FileDescriptorSet message. Does not implicitly {@link google.protobuf.FileDescriptorSet.verify|verify} messages.
             * @param message FileDescriptorSet message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.FileDescriptorSet, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a FileDescriptorSet message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns FileDescriptorSet
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.FileDescriptorSet;
        }

        /** Properties of a FileDescriptorProto. */
        interface IFileDescriptorProto {

            /** FileDescriptorProto name */
            name?: (string|null);

            /** FileDescriptorProto package */
            "package"?: (string|null);

            /** FileDescriptorProto dependency */
            dependency?: (string[]|null);

            /** FileDescriptorProto public_dependency */
            public_dependency?: (number[]|null);

            /** FileDescriptorProto weak_dependency */
            weak_dependency?: (number[]|null);

            /** FileDescriptorProto message_type */
            message_type?: (google.protobuf.DescriptorProto[]|null);

            /** FileDescriptorProto enum_type */
            enum_type?: (google.protobuf.EnumDescriptorProto[]|null);

            /** FileDescriptorProto service */
            service?: (google.protobuf.ServiceDescriptorProto[]|null);

            /** FileDescriptorProto extension */
            extension?: (google.protobuf.FieldDescriptorProto[]|null);

            /** FileDescriptorProto options */
            options?: (google.protobuf.FileOptions|null);

            /** FileDescriptorProto source_code_info */
            source_code_info?: (google.protobuf.SourceCodeInfo|null);

            /** FileDescriptorProto syntax */
            syntax?: (string|null);
        }

        /** Represents a FileDescriptorProto. */
        class FileDescriptorProto implements IFileDescriptorProto {

            /**
             * Constructs a new FileDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IFileDescriptorProto);

            /** FileDescriptorProto name. */
            public name: string;

            /** FileDescriptorProto package. */
            public package: string;

            /** FileDescriptorProto dependency. */
            public dependency: string[];

            /** FileDescriptorProto public_dependency. */
            public public_dependency: number[];

            /** FileDescriptorProto weak_dependency. */
            public weak_dependency: number[];

            /** FileDescriptorProto message_type. */
            public message_type: google.protobuf.DescriptorProto[];

            /** FileDescriptorProto enum_type. */
            public enum_type: google.protobuf.EnumDescriptorProto[];

            /** FileDescriptorProto service. */
            public service: google.protobuf.ServiceDescriptorProto[];

            /** FileDescriptorProto extension. */
            public extension: google.protobuf.FieldDescriptorProto[];

            /** FileDescriptorProto options. */
            public options?: (google.protobuf.FileOptions|null);

            /** FileDescriptorProto source_code_info. */
            public source_code_info?: (google.protobuf.SourceCodeInfo|null);

            /** FileDescriptorProto syntax. */
            public syntax: string;

            /**
             * Encodes the specified FileDescriptorProto message. Does not implicitly {@link google.protobuf.FileDescriptorProto.verify|verify} messages.
             * @param message FileDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.FileDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a FileDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns FileDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.FileDescriptorProto;
        }

        /** Properties of a DescriptorProto. */
        interface IDescriptorProto {

            /** DescriptorProto name */
            name?: (string|null);

            /** DescriptorProto field */
            field?: (google.protobuf.FieldDescriptorProto[]|null);

            /** DescriptorProto extension */
            extension?: (google.protobuf.FieldDescriptorProto[]|null);

            /** DescriptorProto nested_type */
            nested_type?: (google.protobuf.DescriptorProto[]|null);

            /** DescriptorProto enum_type */
            enum_type?: (google.protobuf.EnumDescriptorProto[]|null);

            /** DescriptorProto extension_range */
            extension_range?: (google.protobuf.DescriptorProto.ExtensionRange[]|null);

            /** DescriptorProto oneof_decl */
            oneof_decl?: (google.protobuf.OneofDescriptorProto[]|null);

            /** DescriptorProto options */
            options?: (google.protobuf.MessageOptions|null);

            /** DescriptorProto reserved_range */
            reserved_range?: (google.protobuf.DescriptorProto.ReservedRange[]|null);

            /** DescriptorProto reserved_name */
            reserved_name?: (string[]|null);
        }

        /** Represents a DescriptorProto. */
        class DescriptorProto implements IDescriptorProto {

            /**
             * Constructs a new DescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IDescriptorProto);

            /** DescriptorProto name. */
            public name: string;

            /** DescriptorProto field. */
            public field: google.protobuf.FieldDescriptorProto[];

            /** DescriptorProto extension. */
            public extension: google.protobuf.FieldDescriptorProto[];

            /** DescriptorProto nested_type. */
            public nested_type: google.protobuf.DescriptorProto[];

            /** DescriptorProto enum_type. */
            public enum_type: google.protobuf.EnumDescriptorProto[];

            /** DescriptorProto extension_range. */
            public extension_range: google.protobuf.DescriptorProto.ExtensionRange[];

            /** DescriptorProto oneof_decl. */
            public oneof_decl: google.protobuf.OneofDescriptorProto[];

            /** DescriptorProto options. */
            public options?: (google.protobuf.MessageOptions|null);

            /** DescriptorProto reserved_range. */
            public reserved_range: google.protobuf.DescriptorProto.ReservedRange[];

            /** DescriptorProto reserved_name. */
            public reserved_name: string[];

            /**
             * Encodes the specified DescriptorProto message. Does not implicitly {@link google.protobuf.DescriptorProto.verify|verify} messages.
             * @param message DescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.DescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a DescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns DescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.DescriptorProto;
        }

        namespace DescriptorProto {

            /** Properties of an ExtensionRange. */
            interface IExtensionRange {

                /** ExtensionRange start */
                start?: (number|null);

                /** ExtensionRange end */
                end?: (number|null);
            }

            /** Represents an ExtensionRange. */
            class ExtensionRange implements IExtensionRange {

                /**
                 * Constructs a new ExtensionRange.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: google.protobuf.DescriptorProto.IExtensionRange);

                /** ExtensionRange start. */
                public start: number;

                /** ExtensionRange end. */
                public end: number;

                /**
                 * Encodes the specified ExtensionRange message. Does not implicitly {@link google.protobuf.DescriptorProto.ExtensionRange.verify|verify} messages.
                 * @param message ExtensionRange message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: google.protobuf.DescriptorProto.ExtensionRange, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes an ExtensionRange message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns ExtensionRange
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.DescriptorProto.ExtensionRange;
            }

            /** Properties of a ReservedRange. */
            interface IReservedRange {

                /** ReservedRange start */
                start?: (number|null);

                /** ReservedRange end */
                end?: (number|null);
            }

            /** Represents a ReservedRange. */
            class ReservedRange implements IReservedRange {

                /**
                 * Constructs a new ReservedRange.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: google.protobuf.DescriptorProto.IReservedRange);

                /** ReservedRange start. */
                public start: number;

                /** ReservedRange end. */
                public end: number;

                /**
                 * Encodes the specified ReservedRange message. Does not implicitly {@link google.protobuf.DescriptorProto.ReservedRange.verify|verify} messages.
                 * @param message ReservedRange message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: google.protobuf.DescriptorProto.ReservedRange, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a ReservedRange message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns ReservedRange
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.DescriptorProto.ReservedRange;
            }
        }

        /** Properties of a FieldDescriptorProto. */
        interface IFieldDescriptorProto {

            /** FieldDescriptorProto name */
            name?: (string|null);

            /** FieldDescriptorProto number */
            number?: (number|null);

            /** FieldDescriptorProto label */
            label?: (google.protobuf.FieldDescriptorProto.Label|null);

            /** FieldDescriptorProto type */
            type?: (google.protobuf.FieldDescriptorProto.Type|null);

            /** FieldDescriptorProto type_name */
            type_name?: (string|null);

            /** FieldDescriptorProto extendee */
            extendee?: (string|null);

            /** FieldDescriptorProto default_value */
            default_value?: (string|null);

            /** FieldDescriptorProto oneof_index */
            oneof_index?: (number|null);

            /** FieldDescriptorProto json_name */
            json_name?: (string|null);

            /** FieldDescriptorProto options */
            options?: (google.protobuf.FieldOptions|null);
        }

        /** Represents a FieldDescriptorProto. */
        class FieldDescriptorProto implements IFieldDescriptorProto {

            /**
             * Constructs a new FieldDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IFieldDescriptorProto);

            /** FieldDescriptorProto name. */
            public name: string;

            /** FieldDescriptorProto number. */
            public number: number;

            /** FieldDescriptorProto label. */
            public label: google.protobuf.FieldDescriptorProto.Label;

            /** FieldDescriptorProto type. */
            public type: google.protobuf.FieldDescriptorProto.Type;

            /** FieldDescriptorProto type_name. */
            public type_name: string;

            /** FieldDescriptorProto extendee. */
            public extendee: string;

            /** FieldDescriptorProto default_value. */
            public default_value: string;

            /** FieldDescriptorProto oneof_index. */
            public oneof_index: number;

            /** FieldDescriptorProto json_name. */
            public json_name: string;

            /** FieldDescriptorProto options. */
            public options?: (google.protobuf.FieldOptions|null);

            /**
             * Encodes the specified FieldDescriptorProto message. Does not implicitly {@link google.protobuf.FieldDescriptorProto.verify|verify} messages.
             * @param message FieldDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.FieldDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a FieldDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns FieldDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.FieldDescriptorProto;
        }

        namespace FieldDescriptorProto {

            /** Type enum. */
            enum Type {
                TYPE_DOUBLE = 1,
                TYPE_FLOAT = 2,
                TYPE_INT64 = 3,
                TYPE_UINT64 = 4,
                TYPE_INT32 = 5,
                TYPE_FIXED64 = 6,
                TYPE_FIXED32 = 7,
                TYPE_BOOL = 8,
                TYPE_STRING = 9,
                TYPE_GROUP = 10,
                TYPE_MESSAGE = 11,
                TYPE_BYTES = 12,
                TYPE_UINT32 = 13,
                TYPE_ENUM = 14,
                TYPE_SFIXED32 = 15,
                TYPE_SFIXED64 = 16,
                TYPE_SINT32 = 17,
                TYPE_SINT64 = 18
            }

            /** Label enum. */
            enum Label {
                LABEL_OPTIONAL = 1,
                LABEL_REQUIRED = 2,
                LABEL_REPEATED = 3
            }
        }

        /** Properties of an OneofDescriptorProto. */
        interface IOneofDescriptorProto {

            /** OneofDescriptorProto name */
            name?: (string|null);

            /** OneofDescriptorProto options */
            options?: (google.protobuf.OneofOptions|null);
        }

        /** Represents an OneofDescriptorProto. */
        class OneofDescriptorProto implements IOneofDescriptorProto {

            /**
             * Constructs a new OneofDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IOneofDescriptorProto);

            /** OneofDescriptorProto name. */
            public name: string;

            /** OneofDescriptorProto options. */
            public options?: (google.protobuf.OneofOptions|null);

            /**
             * Encodes the specified OneofDescriptorProto message. Does not implicitly {@link google.protobuf.OneofDescriptorProto.verify|verify} messages.
             * @param message OneofDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.OneofDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an OneofDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns OneofDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.OneofDescriptorProto;
        }

        /** Properties of an EnumDescriptorProto. */
        interface IEnumDescriptorProto {

            /** EnumDescriptorProto name */
            name?: (string|null);

            /** EnumDescriptorProto value */
            value?: (google.protobuf.EnumValueDescriptorProto[]|null);

            /** EnumDescriptorProto options */
            options?: (google.protobuf.EnumOptions|null);
        }

        /** Represents an EnumDescriptorProto. */
        class EnumDescriptorProto implements IEnumDescriptorProto {

            /**
             * Constructs a new EnumDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEnumDescriptorProto);

            /** EnumDescriptorProto name. */
            public name: string;

            /** EnumDescriptorProto value. */
            public value: google.protobuf.EnumValueDescriptorProto[];

            /** EnumDescriptorProto options. */
            public options?: (google.protobuf.EnumOptions|null);

            /**
             * Encodes the specified EnumDescriptorProto message. Does not implicitly {@link google.protobuf.EnumDescriptorProto.verify|verify} messages.
             * @param message EnumDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.EnumDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an EnumDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns EnumDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.EnumDescriptorProto;
        }

        /** Properties of an EnumValueDescriptorProto. */
        interface IEnumValueDescriptorProto {

            /** EnumValueDescriptorProto name */
            name?: (string|null);

            /** EnumValueDescriptorProto number */
            number?: (number|null);

            /** EnumValueDescriptorProto options */
            options?: (google.protobuf.EnumValueOptions|null);
        }

        /** Represents an EnumValueDescriptorProto. */
        class EnumValueDescriptorProto implements IEnumValueDescriptorProto {

            /**
             * Constructs a new EnumValueDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEnumValueDescriptorProto);

            /** EnumValueDescriptorProto name. */
            public name: string;

            /** EnumValueDescriptorProto number. */
            public number: number;

            /** EnumValueDescriptorProto options. */
            public options?: (google.protobuf.EnumValueOptions|null);

            /**
             * Encodes the specified EnumValueDescriptorProto message. Does not implicitly {@link google.protobuf.EnumValueDescriptorProto.verify|verify} messages.
             * @param message EnumValueDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.EnumValueDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an EnumValueDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns EnumValueDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.EnumValueDescriptorProto;
        }

        /** Properties of a ServiceDescriptorProto. */
        interface IServiceDescriptorProto {

            /** ServiceDescriptorProto name */
            name?: (string|null);

            /** ServiceDescriptorProto method */
            method?: (google.protobuf.MethodDescriptorProto[]|null);

            /** ServiceDescriptorProto options */
            options?: (google.protobuf.ServiceOptions|null);
        }

        /** Represents a ServiceDescriptorProto. */
        class ServiceDescriptorProto implements IServiceDescriptorProto {

            /**
             * Constructs a new ServiceDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IServiceDescriptorProto);

            /** ServiceDescriptorProto name. */
            public name: string;

            /** ServiceDescriptorProto method. */
            public method: google.protobuf.MethodDescriptorProto[];

            /** ServiceDescriptorProto options. */
            public options?: (google.protobuf.ServiceOptions|null);

            /**
             * Encodes the specified ServiceDescriptorProto message. Does not implicitly {@link google.protobuf.ServiceDescriptorProto.verify|verify} messages.
             * @param message ServiceDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.ServiceDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a ServiceDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ServiceDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.ServiceDescriptorProto;
        }

        /** Properties of a MethodDescriptorProto. */
        interface IMethodDescriptorProto {

            /** MethodDescriptorProto name */
            name?: (string|null);

            /** MethodDescriptorProto input_type */
            input_type?: (string|null);

            /** MethodDescriptorProto output_type */
            output_type?: (string|null);

            /** MethodDescriptorProto options */
            options?: (google.protobuf.MethodOptions|null);

            /** MethodDescriptorProto client_streaming */
            client_streaming?: (boolean|null);

            /** MethodDescriptorProto server_streaming */
            server_streaming?: (boolean|null);
        }

        /** Represents a MethodDescriptorProto. */
        class MethodDescriptorProto implements IMethodDescriptorProto {

            /**
             * Constructs a new MethodDescriptorProto.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IMethodDescriptorProto);

            /** MethodDescriptorProto name. */
            public name: string;

            /** MethodDescriptorProto input_type. */
            public input_type: string;

            /** MethodDescriptorProto output_type. */
            public output_type: string;

            /** MethodDescriptorProto options. */
            public options?: (google.protobuf.MethodOptions|null);

            /** MethodDescriptorProto client_streaming. */
            public client_streaming: boolean;

            /** MethodDescriptorProto server_streaming. */
            public server_streaming: boolean;

            /**
             * Encodes the specified MethodDescriptorProto message. Does not implicitly {@link google.protobuf.MethodDescriptorProto.verify|verify} messages.
             * @param message MethodDescriptorProto message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.MethodDescriptorProto, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a MethodDescriptorProto message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns MethodDescriptorProto
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.MethodDescriptorProto;
        }

        /** Properties of a FileOptions. */
        interface IFileOptions {

            /** FileOptions java_package */
            java_package?: (string|null);

            /** FileOptions java_outer_classname */
            java_outer_classname?: (string|null);

            /** FileOptions java_multiple_files */
            java_multiple_files?: (boolean|null);

            /** FileOptions java_generate_equals_and_hash */
            java_generate_equals_and_hash?: (boolean|null);

            /** FileOptions java_string_check_utf8 */
            java_string_check_utf8?: (boolean|null);

            /** FileOptions optimize_for */
            optimize_for?: (google.protobuf.FileOptions.OptimizeMode|null);

            /** FileOptions go_package */
            go_package?: (string|null);

            /** FileOptions cc_generic_services */
            cc_generic_services?: (boolean|null);

            /** FileOptions java_generic_services */
            java_generic_services?: (boolean|null);

            /** FileOptions py_generic_services */
            py_generic_services?: (boolean|null);

            /** FileOptions deprecated */
            deprecated?: (boolean|null);

            /** FileOptions cc_enable_arenas */
            cc_enable_arenas?: (boolean|null);

            /** FileOptions objc_class_prefix */
            objc_class_prefix?: (string|null);

            /** FileOptions csharp_namespace */
            csharp_namespace?: (string|null);

            /** FileOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents a FileOptions. */
        class FileOptions implements IFileOptions {

            /**
             * Constructs a new FileOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IFileOptions);

            /** FileOptions java_package. */
            public java_package: string;

            /** FileOptions java_outer_classname. */
            public java_outer_classname: string;

            /** FileOptions java_multiple_files. */
            public java_multiple_files: boolean;

            /** FileOptions java_generate_equals_and_hash. */
            public java_generate_equals_and_hash: boolean;

            /** FileOptions java_string_check_utf8. */
            public java_string_check_utf8: boolean;

            /** FileOptions optimize_for. */
            public optimize_for: google.protobuf.FileOptions.OptimizeMode;

            /** FileOptions go_package. */
            public go_package: string;

            /** FileOptions cc_generic_services. */
            public cc_generic_services: boolean;

            /** FileOptions java_generic_services. */
            public java_generic_services: boolean;

            /** FileOptions py_generic_services. */
            public py_generic_services: boolean;

            /** FileOptions deprecated. */
            public deprecated: boolean;

            /** FileOptions cc_enable_arenas. */
            public cc_enable_arenas: boolean;

            /** FileOptions objc_class_prefix. */
            public objc_class_prefix: string;

            /** FileOptions csharp_namespace. */
            public csharp_namespace: string;

            /** FileOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified FileOptions message. Does not implicitly {@link google.protobuf.FileOptions.verify|verify} messages.
             * @param message FileOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.FileOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a FileOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns FileOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.FileOptions;
        }

        namespace FileOptions {

            /** OptimizeMode enum. */
            enum OptimizeMode {
                SPEED = 1,
                CODE_SIZE = 2,
                LITE_RUNTIME = 3
            }
        }

        /** Properties of a MessageOptions. */
        interface IMessageOptions {

            /** MessageOptions message_set_wire_format */
            message_set_wire_format?: (boolean|null);

            /** MessageOptions no_standard_descriptor_accessor */
            no_standard_descriptor_accessor?: (boolean|null);

            /** MessageOptions deprecated */
            deprecated?: (boolean|null);

            /** MessageOptions map_entry */
            map_entry?: (boolean|null);

            /** MessageOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents a MessageOptions. */
        class MessageOptions implements IMessageOptions {

            /**
             * Constructs a new MessageOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IMessageOptions);

            /** MessageOptions message_set_wire_format. */
            public message_set_wire_format: boolean;

            /** MessageOptions no_standard_descriptor_accessor. */
            public no_standard_descriptor_accessor: boolean;

            /** MessageOptions deprecated. */
            public deprecated: boolean;

            /** MessageOptions map_entry. */
            public map_entry: boolean;

            /** MessageOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified MessageOptions message. Does not implicitly {@link google.protobuf.MessageOptions.verify|verify} messages.
             * @param message MessageOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.MessageOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a MessageOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns MessageOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.MessageOptions;
        }

        /** Properties of a FieldOptions. */
        interface IFieldOptions {

            /** FieldOptions ctype */
            ctype?: (google.protobuf.FieldOptions.CType|null);

            /** FieldOptions packed */
            packed?: (boolean|null);

            /** FieldOptions jstype */
            jstype?: (google.protobuf.FieldOptions.JSType|null);

            /** FieldOptions lazy */
            lazy?: (boolean|null);

            /** FieldOptions deprecated */
            deprecated?: (boolean|null);

            /** FieldOptions weak */
            weak?: (boolean|null);

            /** FieldOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents a FieldOptions. */
        class FieldOptions implements IFieldOptions {

            /**
             * Constructs a new FieldOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IFieldOptions);

            /** FieldOptions ctype. */
            public ctype: google.protobuf.FieldOptions.CType;

            /** FieldOptions packed. */
            public packed: boolean;

            /** FieldOptions jstype. */
            public jstype: google.protobuf.FieldOptions.JSType;

            /** FieldOptions lazy. */
            public lazy: boolean;

            /** FieldOptions deprecated. */
            public deprecated: boolean;

            /** FieldOptions weak. */
            public weak: boolean;

            /** FieldOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified FieldOptions message. Does not implicitly {@link google.protobuf.FieldOptions.verify|verify} messages.
             * @param message FieldOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.FieldOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a FieldOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns FieldOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.FieldOptions;
        }

        namespace FieldOptions {

            /** CType enum. */
            enum CType {
                STRING = 0,
                CORD = 1,
                STRING_PIECE = 2
            }

            /** JSType enum. */
            enum JSType {
                JS_NORMAL = 0,
                JS_STRING = 1,
                JS_NUMBER = 2
            }
        }

        /** Properties of an OneofOptions. */
        interface IOneofOptions {

            /** OneofOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents an OneofOptions. */
        class OneofOptions implements IOneofOptions {

            /**
             * Constructs a new OneofOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IOneofOptions);

            /** OneofOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified OneofOptions message. Does not implicitly {@link google.protobuf.OneofOptions.verify|verify} messages.
             * @param message OneofOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.OneofOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an OneofOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns OneofOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.OneofOptions;
        }

        /** Properties of an EnumOptions. */
        interface IEnumOptions {

            /** EnumOptions allow_alias */
            allow_alias?: (boolean|null);

            /** EnumOptions deprecated */
            deprecated?: (boolean|null);

            /** EnumOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents an EnumOptions. */
        class EnumOptions implements IEnumOptions {

            /**
             * Constructs a new EnumOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEnumOptions);

            /** EnumOptions allow_alias. */
            public allow_alias: boolean;

            /** EnumOptions deprecated. */
            public deprecated: boolean;

            /** EnumOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified EnumOptions message. Does not implicitly {@link google.protobuf.EnumOptions.verify|verify} messages.
             * @param message EnumOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.EnumOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an EnumOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns EnumOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.EnumOptions;
        }

        /** Properties of an EnumValueOptions. */
        interface IEnumValueOptions {

            /** EnumValueOptions deprecated */
            deprecated?: (boolean|null);

            /** EnumValueOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents an EnumValueOptions. */
        class EnumValueOptions implements IEnumValueOptions {

            /**
             * Constructs a new EnumValueOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEnumValueOptions);

            /** EnumValueOptions deprecated. */
            public deprecated: boolean;

            /** EnumValueOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified EnumValueOptions message. Does not implicitly {@link google.protobuf.EnumValueOptions.verify|verify} messages.
             * @param message EnumValueOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.EnumValueOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an EnumValueOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns EnumValueOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.EnumValueOptions;
        }

        /** Properties of a ServiceOptions. */
        interface IServiceOptions {

            /** ServiceOptions deprecated */
            deprecated?: (boolean|null);

            /** ServiceOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);
        }

        /** Represents a ServiceOptions. */
        class ServiceOptions implements IServiceOptions {

            /**
             * Constructs a new ServiceOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IServiceOptions);

            /** ServiceOptions deprecated. */
            public deprecated: boolean;

            /** ServiceOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified ServiceOptions message. Does not implicitly {@link google.protobuf.ServiceOptions.verify|verify} messages.
             * @param message ServiceOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.ServiceOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a ServiceOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns ServiceOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.ServiceOptions;
        }

        /** Properties of a MethodOptions. */
        interface IMethodOptions {

            /** MethodOptions deprecated */
            deprecated?: (boolean|null);

            /** MethodOptions uninterpreted_option */
            uninterpreted_option?: (google.protobuf.UninterpretedOption[]|null);

            /** MethodOptions .google.api.http */
            ".google.api.http"?: (google.api.HttpRule|null);
        }

        /** Represents a MethodOptions. */
        class MethodOptions implements IMethodOptions {

            /**
             * Constructs a new MethodOptions.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IMethodOptions);

            /** MethodOptions deprecated. */
            public deprecated: boolean;

            /** MethodOptions uninterpreted_option. */
            public uninterpreted_option: google.protobuf.UninterpretedOption[];

            /**
             * Encodes the specified MethodOptions message. Does not implicitly {@link google.protobuf.MethodOptions.verify|verify} messages.
             * @param message MethodOptions message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.MethodOptions, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a MethodOptions message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns MethodOptions
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.MethodOptions;
        }

        /** Properties of an UninterpretedOption. */
        interface IUninterpretedOption {

            /** UninterpretedOption name */
            name?: (google.protobuf.UninterpretedOption.NamePart[]|null);

            /** UninterpretedOption identifier_value */
            identifier_value?: (string|null);

            /** UninterpretedOption positive_int_value */
            positive_int_value?: (number|null);

            /** UninterpretedOption negative_int_value */
            negative_int_value?: (number|null);

            /** UninterpretedOption double_value */
            double_value?: (number|null);

            /** UninterpretedOption string_value */
            string_value?: (Uint8Array|null);

            /** UninterpretedOption aggregate_value */
            aggregate_value?: (string|null);
        }

        /** Represents an UninterpretedOption. */
        class UninterpretedOption implements IUninterpretedOption {

            /**
             * Constructs a new UninterpretedOption.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IUninterpretedOption);

            /** UninterpretedOption name. */
            public name: google.protobuf.UninterpretedOption.NamePart[];

            /** UninterpretedOption identifier_value. */
            public identifier_value: string;

            /** UninterpretedOption positive_int_value. */
            public positive_int_value: number;

            /** UninterpretedOption negative_int_value. */
            public negative_int_value: number;

            /** UninterpretedOption double_value. */
            public double_value: number;

            /** UninterpretedOption string_value. */
            public string_value: Uint8Array;

            /** UninterpretedOption aggregate_value. */
            public aggregate_value: string;

            /**
             * Encodes the specified UninterpretedOption message. Does not implicitly {@link google.protobuf.UninterpretedOption.verify|verify} messages.
             * @param message UninterpretedOption message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.UninterpretedOption, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an UninterpretedOption message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns UninterpretedOption
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.UninterpretedOption;
        }

        namespace UninterpretedOption {

            /** Properties of a NamePart. */
            interface INamePart {

                /** NamePart name_part */
                name_part: string;

                /** NamePart is_extension */
                is_extension: boolean;
            }

            /** Represents a NamePart. */
            class NamePart implements INamePart {

                /**
                 * Constructs a new NamePart.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: google.protobuf.UninterpretedOption.INamePart);

                /** NamePart name_part. */
                public name_part: string;

                /** NamePart is_extension. */
                public is_extension: boolean;

                /**
                 * Encodes the specified NamePart message. Does not implicitly {@link google.protobuf.UninterpretedOption.NamePart.verify|verify} messages.
                 * @param message NamePart message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: google.protobuf.UninterpretedOption.NamePart, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a NamePart message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns NamePart
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.UninterpretedOption.NamePart;
            }
        }

        /** Properties of a SourceCodeInfo. */
        interface ISourceCodeInfo {

            /** SourceCodeInfo location */
            location?: (google.protobuf.SourceCodeInfo.Location[]|null);
        }

        /** Represents a SourceCodeInfo. */
        class SourceCodeInfo implements ISourceCodeInfo {

            /**
             * Constructs a new SourceCodeInfo.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.ISourceCodeInfo);

            /** SourceCodeInfo location. */
            public location: google.protobuf.SourceCodeInfo.Location[];

            /**
             * Encodes the specified SourceCodeInfo message. Does not implicitly {@link google.protobuf.SourceCodeInfo.verify|verify} messages.
             * @param message SourceCodeInfo message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.SourceCodeInfo, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a SourceCodeInfo message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns SourceCodeInfo
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.SourceCodeInfo;
        }

        namespace SourceCodeInfo {

            /** Properties of a Location. */
            interface ILocation {

                /** Location path */
                path?: (number[]|null);

                /** Location span */
                span?: (number[]|null);

                /** Location leading_comments */
                leading_comments?: (string|null);

                /** Location trailing_comments */
                trailing_comments?: (string|null);

                /** Location leading_detached_comments */
                leading_detached_comments?: (string[]|null);
            }

            /** Represents a Location. */
            class Location implements ILocation {

                /**
                 * Constructs a new Location.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: google.protobuf.SourceCodeInfo.ILocation);

                /** Location path. */
                public path: number[];

                /** Location span. */
                public span: number[];

                /** Location leading_comments. */
                public leading_comments: string;

                /** Location trailing_comments. */
                public trailing_comments: string;

                /** Location leading_detached_comments. */
                public leading_detached_comments: string[];

                /**
                 * Encodes the specified Location message. Does not implicitly {@link google.protobuf.SourceCodeInfo.Location.verify|verify} messages.
                 * @param message Location message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: google.protobuf.SourceCodeInfo.Location, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a Location message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns Location
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.SourceCodeInfo.Location;
            }
        }

        /** Properties of a GeneratedCodeInfo. */
        interface IGeneratedCodeInfo {

            /** GeneratedCodeInfo annotation */
            annotation?: (google.protobuf.GeneratedCodeInfo.Annotation[]|null);
        }

        /** Represents a GeneratedCodeInfo. */
        class GeneratedCodeInfo implements IGeneratedCodeInfo {

            /**
             * Constructs a new GeneratedCodeInfo.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IGeneratedCodeInfo);

            /** GeneratedCodeInfo annotation. */
            public annotation: google.protobuf.GeneratedCodeInfo.Annotation[];

            /**
             * Encodes the specified GeneratedCodeInfo message. Does not implicitly {@link google.protobuf.GeneratedCodeInfo.verify|verify} messages.
             * @param message GeneratedCodeInfo message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.GeneratedCodeInfo, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a GeneratedCodeInfo message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns GeneratedCodeInfo
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.GeneratedCodeInfo;
        }

        namespace GeneratedCodeInfo {

            /** Properties of an Annotation. */
            interface IAnnotation {

                /** Annotation path */
                path?: (number[]|null);

                /** Annotation source_file */
                source_file?: (string|null);

                /** Annotation begin */
                begin?: (number|null);

                /** Annotation end */
                end?: (number|null);
            }

            /** Represents an Annotation. */
            class Annotation implements IAnnotation {

                /**
                 * Constructs a new Annotation.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: google.protobuf.GeneratedCodeInfo.IAnnotation);

                /** Annotation path. */
                public path: number[];

                /** Annotation source_file. */
                public source_file: string;

                /** Annotation begin. */
                public begin: number;

                /** Annotation end. */
                public end: number;

                /**
                 * Encodes the specified Annotation message. Does not implicitly {@link google.protobuf.GeneratedCodeInfo.Annotation.verify|verify} messages.
                 * @param message Annotation message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: google.protobuf.GeneratedCodeInfo.Annotation, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes an Annotation message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns Annotation
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.GeneratedCodeInfo.Annotation;
            }
        }
    }
}

/** Namespace changelog. */
export namespace changelog {

    /** Properties of a ShowRequest. */
    interface IShowRequest {

        /** ShowRequest project_id */
        project_id?: (number|null);

        /** ShowRequest only_changed */
        only_changed?: (boolean|null);
    }

    /** Represents a ShowRequest. */
    class ShowRequest implements IShowRequest {

        /**
         * Constructs a new ShowRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: changelog.IShowRequest);

        /** ShowRequest project_id. */
        public project_id: number;

        /** ShowRequest only_changed. */
        public only_changed: boolean;

        /**
         * Encodes the specified ShowRequest message. Does not implicitly {@link changelog.ShowRequest.verify|verify} messages.
         * @param message ShowRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: changelog.ShowRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): changelog.ShowRequest;
    }

    /** Properties of a ShowResponse. */
    interface IShowResponse {

        /** ShowResponse items */
        items?: (types.ChangelogModel[]|null);
    }

    /** Represents a ShowResponse. */
    class ShowResponse implements IShowResponse {

        /**
         * Constructs a new ShowResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: changelog.IShowResponse);

        /** ShowResponse items. */
        public items: types.ChangelogModel[];

        /**
         * Encodes the specified ShowResponse message. Does not implicitly {@link changelog.ShowResponse.verify|verify} messages.
         * @param message ShowResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: changelog.ShowResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): changelog.ShowResponse;
    }

    /** Represents a Changelog */
    class Changelog extends $protobuf.rpc.Service {

        /**
         * Constructs a new Changelog service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ShowResponse
         */
        public show(request: changelog.ShowRequest, callback: changelog.Changelog.ShowCallback): void;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @returns Promise
         */
        public show(request: changelog.ShowRequest): Promise<changelog.ShowResponse>;
    }

    namespace Changelog {

        /**
         * Callback as used by {@link changelog.Changelog#show}.
         * @param error Error, if any
         * @param [response] ShowResponse
         */
        type ShowCallback = (error: (Error|null), response?: changelog.ShowResponse) => void;
    }
}

/** Namespace cluster. */
export namespace cluster {

    /** Status enum. */
    enum Status {
        StatusUnknown = 0,
        StatusBad = 1,
        StatusNotGood = 2,
        StatusHealth = 3
    }

    /** Properties of an InfoResponse. */
    interface IInfoResponse {

        /** InfoResponse status */
        status?: (string|null);

        /** InfoResponse free_memory */
        free_memory?: (string|null);

        /** InfoResponse free_cpu */
        free_cpu?: (string|null);

        /** InfoResponse free_request_memory */
        free_request_memory?: (string|null);

        /** InfoResponse free_request_cpu */
        free_request_cpu?: (string|null);

        /** InfoResponse total_memory */
        total_memory?: (string|null);

        /** InfoResponse total_cpu */
        total_cpu?: (string|null);

        /** InfoResponse usage_memory_rate */
        usage_memory_rate?: (string|null);

        /** InfoResponse usage_cpu_rate */
        usage_cpu_rate?: (string|null);

        /** InfoResponse request_memory_rate */
        request_memory_rate?: (string|null);

        /** InfoResponse request_cpu_rate */
        request_cpu_rate?: (string|null);
    }

    /** Represents an InfoResponse. */
    class InfoResponse implements IInfoResponse {

        /**
         * Constructs a new InfoResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: cluster.IInfoResponse);

        /** InfoResponse status. */
        public status: string;

        /** InfoResponse free_memory. */
        public free_memory: string;

        /** InfoResponse free_cpu. */
        public free_cpu: string;

        /** InfoResponse free_request_memory. */
        public free_request_memory: string;

        /** InfoResponse free_request_cpu. */
        public free_request_cpu: string;

        /** InfoResponse total_memory. */
        public total_memory: string;

        /** InfoResponse total_cpu. */
        public total_cpu: string;

        /** InfoResponse usage_memory_rate. */
        public usage_memory_rate: string;

        /** InfoResponse usage_cpu_rate. */
        public usage_cpu_rate: string;

        /** InfoResponse request_memory_rate. */
        public request_memory_rate: string;

        /** InfoResponse request_cpu_rate. */
        public request_cpu_rate: string;

        /**
         * Encodes the specified InfoResponse message. Does not implicitly {@link cluster.InfoResponse.verify|verify} messages.
         * @param message InfoResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: cluster.InfoResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InfoResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InfoResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): cluster.InfoResponse;
    }

    /** Properties of an InfoRequest. */
    interface IInfoRequest {
    }

    /** Represents an InfoRequest. */
    class InfoRequest implements IInfoRequest {

        /**
         * Constructs a new InfoRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: cluster.IInfoRequest);

        /**
         * Encodes the specified InfoRequest message. Does not implicitly {@link cluster.InfoRequest.verify|verify} messages.
         * @param message InfoRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: cluster.InfoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InfoRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InfoRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): cluster.InfoRequest;
    }

    /** Represents a Cluster */
    class Cluster extends $protobuf.rpc.Service {

        /**
         * Constructs a new Cluster service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls ClusterInfo.
         * @param request InfoRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and InfoResponse
         */
        public clusterInfo(request: cluster.InfoRequest, callback: cluster.Cluster.ClusterInfoCallback): void;

        /**
         * Calls ClusterInfo.
         * @param request InfoRequest message or plain object
         * @returns Promise
         */
        public clusterInfo(request: cluster.InfoRequest): Promise<cluster.InfoResponse>;
    }

    namespace Cluster {

        /**
         * Callback as used by {@link cluster.Cluster#clusterInfo}.
         * @param error Error, if any
         * @param [response] InfoResponse
         */
        type ClusterInfoCallback = (error: (Error|null), response?: cluster.InfoResponse) => void;
    }
}

/** Namespace container. */
export namespace container {

    /** Properties of a CopyToPodRequest. */
    interface ICopyToPodRequest {

        /** CopyToPodRequest file_id */
        file_id?: (number|null);

        /** CopyToPodRequest namespace */
        namespace?: (string|null);

        /** CopyToPodRequest pod */
        pod?: (string|null);

        /** CopyToPodRequest container */
        container?: (string|null);
    }

    /** Represents a CopyToPodRequest. */
    class CopyToPodRequest implements ICopyToPodRequest {

        /**
         * Constructs a new CopyToPodRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.ICopyToPodRequest);

        /** CopyToPodRequest file_id. */
        public file_id: number;

        /** CopyToPodRequest namespace. */
        public namespace: string;

        /** CopyToPodRequest pod. */
        public pod: string;

        /** CopyToPodRequest container. */
        public container: string;

        /**
         * Encodes the specified CopyToPodRequest message. Does not implicitly {@link container.CopyToPodRequest.verify|verify} messages.
         * @param message CopyToPodRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.CopyToPodRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CopyToPodRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CopyToPodRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.CopyToPodRequest;
    }

    /** Properties of a CopyToPodResponse. */
    interface ICopyToPodResponse {

        /** CopyToPodResponse pod_file_path */
        pod_file_path?: (string|null);

        /** CopyToPodResponse output */
        output?: (string|null);

        /** CopyToPodResponse file_name */
        file_name?: (string|null);
    }

    /** Represents a CopyToPodResponse. */
    class CopyToPodResponse implements ICopyToPodResponse {

        /**
         * Constructs a new CopyToPodResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.ICopyToPodResponse);

        /** CopyToPodResponse pod_file_path. */
        public pod_file_path: string;

        /** CopyToPodResponse output. */
        public output: string;

        /** CopyToPodResponse file_name. */
        public file_name: string;

        /**
         * Encodes the specified CopyToPodResponse message. Does not implicitly {@link container.CopyToPodResponse.verify|verify} messages.
         * @param message CopyToPodResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.CopyToPodResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CopyToPodResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CopyToPodResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.CopyToPodResponse;
    }

    /** Properties of an ExecRequest. */
    interface IExecRequest {

        /** ExecRequest namespace */
        namespace?: (string|null);

        /** ExecRequest pod */
        pod?: (string|null);

        /** ExecRequest container */
        container?: (string|null);

        /** ExecRequest command */
        command?: (string[]|null);
    }

    /** Represents an ExecRequest. */
    class ExecRequest implements IExecRequest {

        /**
         * Constructs a new ExecRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IExecRequest);

        /** ExecRequest namespace. */
        public namespace: string;

        /** ExecRequest pod. */
        public pod: string;

        /** ExecRequest container. */
        public container: string;

        /** ExecRequest command. */
        public command: string[];

        /**
         * Encodes the specified ExecRequest message. Does not implicitly {@link container.ExecRequest.verify|verify} messages.
         * @param message ExecRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.ExecRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExecRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExecRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.ExecRequest;
    }

    /** Properties of an ExecError. */
    interface IExecError {

        /** ExecError code */
        code?: (number|null);

        /** ExecError message */
        message?: (string|null);
    }

    /** Represents an ExecError. */
    class ExecError implements IExecError {

        /**
         * Constructs a new ExecError.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IExecError);

        /** ExecError code. */
        public code: number;

        /** ExecError message. */
        public message: string;

        /**
         * Encodes the specified ExecError message. Does not implicitly {@link container.ExecError.verify|verify} messages.
         * @param message ExecError message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.ExecError, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExecError message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExecError
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.ExecError;
    }

    /** Properties of an ExecResponse. */
    interface IExecResponse {

        /** ExecResponse message */
        message?: (string|null);

        /** ExecResponse error */
        error?: (container.ExecError|null);
    }

    /** Represents an ExecResponse. */
    class ExecResponse implements IExecResponse {

        /**
         * Constructs a new ExecResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IExecResponse);

        /** ExecResponse message. */
        public message: string;

        /** ExecResponse error. */
        public error?: (container.ExecError|null);

        /**
         * Encodes the specified ExecResponse message. Does not implicitly {@link container.ExecResponse.verify|verify} messages.
         * @param message ExecResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.ExecResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExecResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExecResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.ExecResponse;
    }

    /** Properties of a StreamCopyToPodRequest. */
    interface IStreamCopyToPodRequest {

        /** StreamCopyToPodRequest file_name */
        file_name?: (string|null);

        /** StreamCopyToPodRequest data */
        data?: (Uint8Array|null);

        /** StreamCopyToPodRequest namespace */
        namespace?: (string|null);

        /** StreamCopyToPodRequest pod */
        pod?: (string|null);

        /** StreamCopyToPodRequest container */
        container?: (string|null);
    }

    /** Represents a StreamCopyToPodRequest. */
    class StreamCopyToPodRequest implements IStreamCopyToPodRequest {

        /**
         * Constructs a new StreamCopyToPodRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IStreamCopyToPodRequest);

        /** StreamCopyToPodRequest file_name. */
        public file_name: string;

        /** StreamCopyToPodRequest data. */
        public data: Uint8Array;

        /** StreamCopyToPodRequest namespace. */
        public namespace: string;

        /** StreamCopyToPodRequest pod. */
        public pod: string;

        /** StreamCopyToPodRequest container. */
        public container: string;

        /**
         * Encodes the specified StreamCopyToPodRequest message. Does not implicitly {@link container.StreamCopyToPodRequest.verify|verify} messages.
         * @param message StreamCopyToPodRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.StreamCopyToPodRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StreamCopyToPodRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns StreamCopyToPodRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.StreamCopyToPodRequest;
    }

    /** Properties of a StreamCopyToPodResponse. */
    interface IStreamCopyToPodResponse {

        /** StreamCopyToPodResponse size */
        size?: (number|null);

        /** StreamCopyToPodResponse pod_file_path */
        pod_file_path?: (string|null);

        /** StreamCopyToPodResponse output */
        output?: (string|null);

        /** StreamCopyToPodResponse pod */
        pod?: (string|null);

        /** StreamCopyToPodResponse namespace */
        namespace?: (string|null);

        /** StreamCopyToPodResponse container */
        container?: (string|null);

        /** StreamCopyToPodResponse filename */
        filename?: (string|null);
    }

    /** Represents a StreamCopyToPodResponse. */
    class StreamCopyToPodResponse implements IStreamCopyToPodResponse {

        /**
         * Constructs a new StreamCopyToPodResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IStreamCopyToPodResponse);

        /** StreamCopyToPodResponse size. */
        public size: number;

        /** StreamCopyToPodResponse pod_file_path. */
        public pod_file_path: string;

        /** StreamCopyToPodResponse output. */
        public output: string;

        /** StreamCopyToPodResponse pod. */
        public pod: string;

        /** StreamCopyToPodResponse namespace. */
        public namespace: string;

        /** StreamCopyToPodResponse container. */
        public container: string;

        /** StreamCopyToPodResponse filename. */
        public filename: string;

        /**
         * Encodes the specified StreamCopyToPodResponse message. Does not implicitly {@link container.StreamCopyToPodResponse.verify|verify} messages.
         * @param message StreamCopyToPodResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.StreamCopyToPodResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a StreamCopyToPodResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns StreamCopyToPodResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.StreamCopyToPodResponse;
    }

    /** Properties of an IsPodRunningRequest. */
    interface IIsPodRunningRequest {

        /** IsPodRunningRequest namespace */
        namespace?: (string|null);

        /** IsPodRunningRequest pod */
        pod?: (string|null);
    }

    /** Represents an IsPodRunningRequest. */
    class IsPodRunningRequest implements IIsPodRunningRequest {

        /**
         * Constructs a new IsPodRunningRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IIsPodRunningRequest);

        /** IsPodRunningRequest namespace. */
        public namespace: string;

        /** IsPodRunningRequest pod. */
        public pod: string;

        /**
         * Encodes the specified IsPodRunningRequest message. Does not implicitly {@link container.IsPodRunningRequest.verify|verify} messages.
         * @param message IsPodRunningRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.IsPodRunningRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsPodRunningRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsPodRunningRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.IsPodRunningRequest;
    }

    /** Properties of an IsPodRunningResponse. */
    interface IIsPodRunningResponse {

        /** IsPodRunningResponse running */
        running?: (boolean|null);

        /** IsPodRunningResponse reason */
        reason?: (string|null);
    }

    /** Represents an IsPodRunningResponse. */
    class IsPodRunningResponse implements IIsPodRunningResponse {

        /**
         * Constructs a new IsPodRunningResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IIsPodRunningResponse);

        /** IsPodRunningResponse running. */
        public running: boolean;

        /** IsPodRunningResponse reason. */
        public reason: string;

        /**
         * Encodes the specified IsPodRunningResponse message. Does not implicitly {@link container.IsPodRunningResponse.verify|verify} messages.
         * @param message IsPodRunningResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.IsPodRunningResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsPodRunningResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsPodRunningResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.IsPodRunningResponse;
    }

    /** Properties of an IsPodExistsRequest. */
    interface IIsPodExistsRequest {

        /** IsPodExistsRequest namespace */
        namespace?: (string|null);

        /** IsPodExistsRequest pod */
        pod?: (string|null);
    }

    /** Represents an IsPodExistsRequest. */
    class IsPodExistsRequest implements IIsPodExistsRequest {

        /**
         * Constructs a new IsPodExistsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IIsPodExistsRequest);

        /** IsPodExistsRequest namespace. */
        public namespace: string;

        /** IsPodExistsRequest pod. */
        public pod: string;

        /**
         * Encodes the specified IsPodExistsRequest message. Does not implicitly {@link container.IsPodExistsRequest.verify|verify} messages.
         * @param message IsPodExistsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.IsPodExistsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsPodExistsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsPodExistsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.IsPodExistsRequest;
    }

    /** Properties of an IsPodExistsResponse. */
    interface IIsPodExistsResponse {

        /** IsPodExistsResponse exists */
        exists?: (boolean|null);
    }

    /** Represents an IsPodExistsResponse. */
    class IsPodExistsResponse implements IIsPodExistsResponse {

        /**
         * Constructs a new IsPodExistsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.IIsPodExistsResponse);

        /** IsPodExistsResponse exists. */
        public exists: boolean;

        /**
         * Encodes the specified IsPodExistsResponse message. Does not implicitly {@link container.IsPodExistsResponse.verify|verify} messages.
         * @param message IsPodExistsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.IsPodExistsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsPodExistsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsPodExistsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.IsPodExistsResponse;
    }

    /** Properties of a LogRequest. */
    interface ILogRequest {

        /** LogRequest namespace */
        namespace?: (string|null);

        /** LogRequest pod */
        pod?: (string|null);

        /** LogRequest container */
        container?: (string|null);
    }

    /** Represents a LogRequest. */
    class LogRequest implements ILogRequest {

        /**
         * Constructs a new LogRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.ILogRequest);

        /** LogRequest namespace. */
        public namespace: string;

        /** LogRequest pod. */
        public pod: string;

        /** LogRequest container. */
        public container: string;

        /**
         * Encodes the specified LogRequest message. Does not implicitly {@link container.LogRequest.verify|verify} messages.
         * @param message LogRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.LogRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a LogRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns LogRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.LogRequest;
    }

    /** Properties of a LogResponse. */
    interface ILogResponse {

        /** LogResponse namespace */
        namespace?: (string|null);

        /** LogResponse pod_name */
        pod_name?: (string|null);

        /** LogResponse container_name */
        container_name?: (string|null);

        /** LogResponse log */
        log?: (string|null);
    }

    /** Represents a LogResponse. */
    class LogResponse implements ILogResponse {

        /**
         * Constructs a new LogResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: container.ILogResponse);

        /** LogResponse namespace. */
        public namespace: string;

        /** LogResponse pod_name. */
        public pod_name: string;

        /** LogResponse container_name. */
        public container_name: string;

        /** LogResponse log. */
        public log: string;

        /**
         * Encodes the specified LogResponse message. Does not implicitly {@link container.LogResponse.verify|verify} messages.
         * @param message LogResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: container.LogResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a LogResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns LogResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): container.LogResponse;
    }

    /** Represents a Container */
    class Container extends $protobuf.rpc.Service {

        /**
         * Constructs a new Container service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls CopyToPod.
         * @param request CopyToPodRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CopyToPodResponse
         */
        public copyToPod(request: container.CopyToPodRequest, callback: container.Container.CopyToPodCallback): void;

        /**
         * Calls CopyToPod.
         * @param request CopyToPodRequest message or plain object
         * @returns Promise
         */
        public copyToPod(request: container.CopyToPodRequest): Promise<container.CopyToPodResponse>;

        /**
         * Calls Exec.
         * @param request ExecRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ExecResponse
         */
        public exec(request: container.ExecRequest, callback: container.Container.ExecCallback): void;

        /**
         * Calls Exec.
         * @param request ExecRequest message or plain object
         * @returns Promise
         */
        public exec(request: container.ExecRequest): Promise<container.ExecResponse>;

        /**
         * Calls StreamCopyToPod.
         * @param request StreamCopyToPodRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and StreamCopyToPodResponse
         */
        public streamCopyToPod(request: container.StreamCopyToPodRequest, callback: container.Container.StreamCopyToPodCallback): void;

        /**
         * Calls StreamCopyToPod.
         * @param request StreamCopyToPodRequest message or plain object
         * @returns Promise
         */
        public streamCopyToPod(request: container.StreamCopyToPodRequest): Promise<container.StreamCopyToPodResponse>;

        /**
         * Calls IsPodRunning.
         * @param request IsPodRunningRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and IsPodRunningResponse
         */
        public isPodRunning(request: container.IsPodRunningRequest, callback: container.Container.IsPodRunningCallback): void;

        /**
         * Calls IsPodRunning.
         * @param request IsPodRunningRequest message or plain object
         * @returns Promise
         */
        public isPodRunning(request: container.IsPodRunningRequest): Promise<container.IsPodRunningResponse>;

        /**
         * Calls IsPodExists.
         * @param request IsPodExistsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and IsPodExistsResponse
         */
        public isPodExists(request: container.IsPodExistsRequest, callback: container.Container.IsPodExistsCallback): void;

        /**
         * Calls IsPodExists.
         * @param request IsPodExistsRequest message or plain object
         * @returns Promise
         */
        public isPodExists(request: container.IsPodExistsRequest): Promise<container.IsPodExistsResponse>;

        /**
         * Calls ContainerLog.
         * @param request LogRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and LogResponse
         */
        public containerLog(request: container.LogRequest, callback: container.Container.ContainerLogCallback): void;

        /**
         * Calls ContainerLog.
         * @param request LogRequest message or plain object
         * @returns Promise
         */
        public containerLog(request: container.LogRequest): Promise<container.LogResponse>;

        /**
         * Calls StreamContainerLog.
         * @param request LogRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and LogResponse
         */
        public streamContainerLog(request: container.LogRequest, callback: container.Container.StreamContainerLogCallback): void;

        /**
         * Calls StreamContainerLog.
         * @param request LogRequest message or plain object
         * @returns Promise
         */
        public streamContainerLog(request: container.LogRequest): Promise<container.LogResponse>;
    }

    namespace Container {

        /**
         * Callback as used by {@link container.Container#copyToPod}.
         * @param error Error, if any
         * @param [response] CopyToPodResponse
         */
        type CopyToPodCallback = (error: (Error|null), response?: container.CopyToPodResponse) => void;

        /**
         * Callback as used by {@link container.Container#exec}.
         * @param error Error, if any
         * @param [response] ExecResponse
         */
        type ExecCallback = (error: (Error|null), response?: container.ExecResponse) => void;

        /**
         * Callback as used by {@link container.Container#streamCopyToPod}.
         * @param error Error, if any
         * @param [response] StreamCopyToPodResponse
         */
        type StreamCopyToPodCallback = (error: (Error|null), response?: container.StreamCopyToPodResponse) => void;

        /**
         * Callback as used by {@link container.Container#isPodRunning}.
         * @param error Error, if any
         * @param [response] IsPodRunningResponse
         */
        type IsPodRunningCallback = (error: (Error|null), response?: container.IsPodRunningResponse) => void;

        /**
         * Callback as used by {@link container.Container#isPodExists}.
         * @param error Error, if any
         * @param [response] IsPodExistsResponse
         */
        type IsPodExistsCallback = (error: (Error|null), response?: container.IsPodExistsResponse) => void;

        /**
         * Callback as used by {@link container.Container#containerLog}.
         * @param error Error, if any
         * @param [response] LogResponse
         */
        type ContainerLogCallback = (error: (Error|null), response?: container.LogResponse) => void;

        /**
         * Callback as used by {@link container.Container#streamContainerLog}.
         * @param error Error, if any
         * @param [response] LogResponse
         */
        type StreamContainerLogCallback = (error: (Error|null), response?: container.LogResponse) => void;
    }
}

/** Namespace endpoint. */
export namespace endpoint {

    /** Properties of an InNamespaceRequest. */
    interface IInNamespaceRequest {

        /** InNamespaceRequest namespace_id */
        namespace_id?: (number|null);
    }

    /** Represents an InNamespaceRequest. */
    class InNamespaceRequest implements IInNamespaceRequest {

        /**
         * Constructs a new InNamespaceRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: endpoint.IInNamespaceRequest);

        /** InNamespaceRequest namespace_id. */
        public namespace_id: number;

        /**
         * Encodes the specified InNamespaceRequest message. Does not implicitly {@link endpoint.InNamespaceRequest.verify|verify} messages.
         * @param message InNamespaceRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: endpoint.InNamespaceRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InNamespaceRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InNamespaceRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): endpoint.InNamespaceRequest;
    }

    /** Properties of an InNamespaceResponse. */
    interface IInNamespaceResponse {

        /** InNamespaceResponse items */
        items?: (types.ServiceEndpoint[]|null);
    }

    /** Represents an InNamespaceResponse. */
    class InNamespaceResponse implements IInNamespaceResponse {

        /**
         * Constructs a new InNamespaceResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: endpoint.IInNamespaceResponse);

        /** InNamespaceResponse items. */
        public items: types.ServiceEndpoint[];

        /**
         * Encodes the specified InNamespaceResponse message. Does not implicitly {@link endpoint.InNamespaceResponse.verify|verify} messages.
         * @param message InNamespaceResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: endpoint.InNamespaceResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InNamespaceResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InNamespaceResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): endpoint.InNamespaceResponse;
    }

    /** Properties of an InProjectRequest. */
    interface IInProjectRequest {

        /** InProjectRequest project_id */
        project_id?: (number|null);
    }

    /** Represents an InProjectRequest. */
    class InProjectRequest implements IInProjectRequest {

        /**
         * Constructs a new InProjectRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: endpoint.IInProjectRequest);

        /** InProjectRequest project_id. */
        public project_id: number;

        /**
         * Encodes the specified InProjectRequest message. Does not implicitly {@link endpoint.InProjectRequest.verify|verify} messages.
         * @param message InProjectRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: endpoint.InProjectRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InProjectRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InProjectRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): endpoint.InProjectRequest;
    }

    /** Properties of an InProjectResponse. */
    interface IInProjectResponse {

        /** InProjectResponse items */
        items?: (types.ServiceEndpoint[]|null);
    }

    /** Represents an InProjectResponse. */
    class InProjectResponse implements IInProjectResponse {

        /**
         * Constructs a new InProjectResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: endpoint.IInProjectResponse);

        /** InProjectResponse items. */
        public items: types.ServiceEndpoint[];

        /**
         * Encodes the specified InProjectResponse message. Does not implicitly {@link endpoint.InProjectResponse.verify|verify} messages.
         * @param message InProjectResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: endpoint.InProjectResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an InProjectResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns InProjectResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): endpoint.InProjectResponse;
    }

    /** Represents an Endpoint */
    class Endpoint extends $protobuf.rpc.Service {

        /**
         * Constructs a new Endpoint service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls InNamespace.
         * @param request InNamespaceRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and InNamespaceResponse
         */
        public inNamespace(request: endpoint.InNamespaceRequest, callback: endpoint.Endpoint.InNamespaceCallback): void;

        /**
         * Calls InNamespace.
         * @param request InNamespaceRequest message or plain object
         * @returns Promise
         */
        public inNamespace(request: endpoint.InNamespaceRequest): Promise<endpoint.InNamespaceResponse>;

        /**
         * Calls InProject.
         * @param request InProjectRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and InProjectResponse
         */
        public inProject(request: endpoint.InProjectRequest, callback: endpoint.Endpoint.InProjectCallback): void;

        /**
         * Calls InProject.
         * @param request InProjectRequest message or plain object
         * @returns Promise
         */
        public inProject(request: endpoint.InProjectRequest): Promise<endpoint.InProjectResponse>;
    }

    namespace Endpoint {

        /**
         * Callback as used by {@link endpoint.Endpoint#inNamespace}.
         * @param error Error, if any
         * @param [response] InNamespaceResponse
         */
        type InNamespaceCallback = (error: (Error|null), response?: endpoint.InNamespaceResponse) => void;

        /**
         * Callback as used by {@link endpoint.Endpoint#inProject}.
         * @param error Error, if any
         * @param [response] InProjectResponse
         */
        type InProjectCallback = (error: (Error|null), response?: endpoint.InProjectResponse) => void;
    }
}

/** Namespace event. */
export namespace event {

    /** Properties of a ListRequest. */
    interface IListRequest {

        /** ListRequest page */
        page?: (number|null);

        /** ListRequest page_size */
        page_size?: (number|null);

        /** ListRequest action_type */
        action_type?: (types.EventActionType|null);
    }

    /** Represents a ListRequest. */
    class ListRequest implements IListRequest {

        /**
         * Constructs a new ListRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: event.IListRequest);

        /** ListRequest page. */
        public page: number;

        /** ListRequest page_size. */
        public page_size: number;

        /** ListRequest action_type. */
        public action_type: types.EventActionType;

        /**
         * Encodes the specified ListRequest message. Does not implicitly {@link event.ListRequest.verify|verify} messages.
         * @param message ListRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: event.ListRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): event.ListRequest;
    }

    /** Properties of a ListResponse. */
    interface IListResponse {

        /** ListResponse page */
        page?: (number|null);

        /** ListResponse page_size */
        page_size?: (number|null);

        /** ListResponse items */
        items?: (types.EventModel[]|null);

        /** ListResponse count */
        count?: (number|null);
    }

    /** Represents a ListResponse. */
    class ListResponse implements IListResponse {

        /**
         * Constructs a new ListResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: event.IListResponse);

        /** ListResponse page. */
        public page: number;

        /** ListResponse page_size. */
        public page_size: number;

        /** ListResponse items. */
        public items: types.EventModel[];

        /** ListResponse count. */
        public count: number;

        /**
         * Encodes the specified ListResponse message. Does not implicitly {@link event.ListResponse.verify|verify} messages.
         * @param message ListResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: event.ListResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): event.ListResponse;
    }

    /** Represents an Event */
    class Event extends $protobuf.rpc.Service {

        /**
         * Constructs a new Event service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ListResponse
         */
        public list(request: event.ListRequest, callback: event.Event.ListCallback): void;

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @returns Promise
         */
        public list(request: event.ListRequest): Promise<event.ListResponse>;
    }

    namespace Event {

        /**
         * Callback as used by {@link event.Event#list}.
         * @param error Error, if any
         * @param [response] ListResponse
         */
        type ListCallback = (error: (Error|null), response?: event.ListResponse) => void;
    }
}

/** Namespace file. */
export namespace file {

    /** Properties of a DeleteRequest. */
    interface IDeleteRequest {

        /** DeleteRequest id */
        id?: (number|null);
    }

    /** Represents a DeleteRequest. */
    class DeleteRequest implements IDeleteRequest {

        /**
         * Constructs a new DeleteRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDeleteRequest);

        /** DeleteRequest id. */
        public id: number;

        /**
         * Encodes the specified DeleteRequest message. Does not implicitly {@link file.DeleteRequest.verify|verify} messages.
         * @param message DeleteRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DeleteRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DeleteRequest;
    }

    /** Properties of a DeleteResponse. */
    interface IDeleteResponse {

        /** DeleteResponse file */
        file?: (types.FileModel|null);
    }

    /** Represents a DeleteResponse. */
    class DeleteResponse implements IDeleteResponse {

        /**
         * Constructs a new DeleteResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDeleteResponse);

        /** DeleteResponse file. */
        public file?: (types.FileModel|null);

        /**
         * Encodes the specified DeleteResponse message. Does not implicitly {@link file.DeleteResponse.verify|verify} messages.
         * @param message DeleteResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DeleteResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DeleteResponse;
    }

    /** Properties of a DeleteUndocumentedFilesRequest. */
    interface IDeleteUndocumentedFilesRequest {
    }

    /** Represents a DeleteUndocumentedFilesRequest. */
    class DeleteUndocumentedFilesRequest implements IDeleteUndocumentedFilesRequest {

        /**
         * Constructs a new DeleteUndocumentedFilesRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDeleteUndocumentedFilesRequest);

        /**
         * Encodes the specified DeleteUndocumentedFilesRequest message. Does not implicitly {@link file.DeleteUndocumentedFilesRequest.verify|verify} messages.
         * @param message DeleteUndocumentedFilesRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DeleteUndocumentedFilesRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteUndocumentedFilesRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteUndocumentedFilesRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DeleteUndocumentedFilesRequest;
    }

    /** Properties of a DeleteUndocumentedFilesResponse. */
    interface IDeleteUndocumentedFilesResponse {

        /** DeleteUndocumentedFilesResponse items */
        items?: (types.FileModel[]|null);
    }

    /** Represents a DeleteUndocumentedFilesResponse. */
    class DeleteUndocumentedFilesResponse implements IDeleteUndocumentedFilesResponse {

        /**
         * Constructs a new DeleteUndocumentedFilesResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDeleteUndocumentedFilesResponse);

        /** DeleteUndocumentedFilesResponse items. */
        public items: types.FileModel[];

        /**
         * Encodes the specified DeleteUndocumentedFilesResponse message. Does not implicitly {@link file.DeleteUndocumentedFilesResponse.verify|verify} messages.
         * @param message DeleteUndocumentedFilesResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DeleteUndocumentedFilesResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteUndocumentedFilesResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteUndocumentedFilesResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DeleteUndocumentedFilesResponse;
    }

    /** Properties of a DiskInfoRequest. */
    interface IDiskInfoRequest {
    }

    /** Represents a DiskInfoRequest. */
    class DiskInfoRequest implements IDiskInfoRequest {

        /**
         * Constructs a new DiskInfoRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDiskInfoRequest);

        /**
         * Encodes the specified DiskInfoRequest message. Does not implicitly {@link file.DiskInfoRequest.verify|verify} messages.
         * @param message DiskInfoRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DiskInfoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DiskInfoRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DiskInfoRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DiskInfoRequest;
    }

    /** Properties of a DiskInfoResponse. */
    interface IDiskInfoResponse {

        /** DiskInfoResponse usage */
        usage?: (number|null);

        /** DiskInfoResponse humanize_usage */
        humanize_usage?: (string|null);
    }

    /** Represents a DiskInfoResponse. */
    class DiskInfoResponse implements IDiskInfoResponse {

        /**
         * Constructs a new DiskInfoResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IDiskInfoResponse);

        /** DiskInfoResponse usage. */
        public usage: number;

        /** DiskInfoResponse humanize_usage. */
        public humanize_usage: string;

        /**
         * Encodes the specified DiskInfoResponse message. Does not implicitly {@link file.DiskInfoResponse.verify|verify} messages.
         * @param message DiskInfoResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.DiskInfoResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DiskInfoResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DiskInfoResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.DiskInfoResponse;
    }

    /** Properties of a ListRequest. */
    interface IListRequest {

        /** ListRequest page */
        page?: (number|null);

        /** ListRequest page_size */
        page_size?: (number|null);

        /** ListRequest without_deleted */
        without_deleted?: (boolean|null);
    }

    /** Represents a ListRequest. */
    class ListRequest implements IListRequest {

        /**
         * Constructs a new ListRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IListRequest);

        /** ListRequest page. */
        public page: number;

        /** ListRequest page_size. */
        public page_size: number;

        /** ListRequest without_deleted. */
        public without_deleted: boolean;

        /**
         * Encodes the specified ListRequest message. Does not implicitly {@link file.ListRequest.verify|verify} messages.
         * @param message ListRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.ListRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.ListRequest;
    }

    /** Properties of a ListResponse. */
    interface IListResponse {

        /** ListResponse page */
        page?: (number|null);

        /** ListResponse page_size */
        page_size?: (number|null);

        /** ListResponse items */
        items?: (types.FileModel[]|null);

        /** ListResponse count */
        count?: (number|null);
    }

    /** Represents a ListResponse. */
    class ListResponse implements IListResponse {

        /**
         * Constructs a new ListResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: file.IListResponse);

        /** ListResponse page. */
        public page: number;

        /** ListResponse page_size. */
        public page_size: number;

        /** ListResponse items. */
        public items: types.FileModel[];

        /** ListResponse count. */
        public count: number;

        /**
         * Encodes the specified ListResponse message. Does not implicitly {@link file.ListResponse.verify|verify} messages.
         * @param message ListResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: file.ListResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): file.ListResponse;
    }

    /** Represents a File */
    class File extends $protobuf.rpc.Service {

        /**
         * Constructs a new File service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ListResponse
         */
        public list(request: file.ListRequest, callback: file.File.ListCallback): void;

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @returns Promise
         */
        public list(request: file.ListRequest): Promise<file.ListResponse>;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DeleteResponse
         */
        public delete(request: file.DeleteRequest, callback: file.File.DeleteCallback): void;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @returns Promise
         */
        public delete(request: file.DeleteRequest): Promise<file.DeleteResponse>;

        /**
         * Calls DeleteUndocumentedFiles.
         * @param request DeleteUndocumentedFilesRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DeleteUndocumentedFilesResponse
         */
        public deleteUndocumentedFiles(request: file.DeleteUndocumentedFilesRequest, callback: file.File.DeleteUndocumentedFilesCallback): void;

        /**
         * Calls DeleteUndocumentedFiles.
         * @param request DeleteUndocumentedFilesRequest message or plain object
         * @returns Promise
         */
        public deleteUndocumentedFiles(request: file.DeleteUndocumentedFilesRequest): Promise<file.DeleteUndocumentedFilesResponse>;

        /**
         * Calls DiskInfo.
         * @param request DiskInfoRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DiskInfoResponse
         */
        public diskInfo(request: file.DiskInfoRequest, callback: file.File.DiskInfoCallback): void;

        /**
         * Calls DiskInfo.
         * @param request DiskInfoRequest message or plain object
         * @returns Promise
         */
        public diskInfo(request: file.DiskInfoRequest): Promise<file.DiskInfoResponse>;
    }

    namespace File {

        /**
         * Callback as used by {@link file.File#list}.
         * @param error Error, if any
         * @param [response] ListResponse
         */
        type ListCallback = (error: (Error|null), response?: file.ListResponse) => void;

        /**
         * Callback as used by {@link file.File#delete_}.
         * @param error Error, if any
         * @param [response] DeleteResponse
         */
        type DeleteCallback = (error: (Error|null), response?: file.DeleteResponse) => void;

        /**
         * Callback as used by {@link file.File#deleteUndocumentedFiles}.
         * @param error Error, if any
         * @param [response] DeleteUndocumentedFilesResponse
         */
        type DeleteUndocumentedFilesCallback = (error: (Error|null), response?: file.DeleteUndocumentedFilesResponse) => void;

        /**
         * Callback as used by {@link file.File#diskInfo}.
         * @param error Error, if any
         * @param [response] DiskInfoResponse
         */
        type DiskInfoCallback = (error: (Error|null), response?: file.DiskInfoResponse) => void;
    }
}

/** Namespace git. */
export namespace git {

    /** Properties of an EnableProjectRequest. */
    interface IEnableProjectRequest {

        /** EnableProjectRequest git_project_id */
        git_project_id?: (string|null);
    }

    /** Represents an EnableProjectRequest. */
    class EnableProjectRequest implements IEnableProjectRequest {

        /**
         * Constructs a new EnableProjectRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IEnableProjectRequest);

        /** EnableProjectRequest git_project_id. */
        public git_project_id: string;

        /**
         * Encodes the specified EnableProjectRequest message. Does not implicitly {@link git.EnableProjectRequest.verify|verify} messages.
         * @param message EnableProjectRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.EnableProjectRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EnableProjectRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns EnableProjectRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.EnableProjectRequest;
    }

    /** Properties of a DisableProjectRequest. */
    interface IDisableProjectRequest {

        /** DisableProjectRequest git_project_id */
        git_project_id?: (string|null);
    }

    /** Represents a DisableProjectRequest. */
    class DisableProjectRequest implements IDisableProjectRequest {

        /**
         * Constructs a new DisableProjectRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IDisableProjectRequest);

        /** DisableProjectRequest git_project_id. */
        public git_project_id: string;

        /**
         * Encodes the specified DisableProjectRequest message. Does not implicitly {@link git.DisableProjectRequest.verify|verify} messages.
         * @param message DisableProjectRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.DisableProjectRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DisableProjectRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DisableProjectRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.DisableProjectRequest;
    }

    /** Properties of a ProjectItem. */
    interface IProjectItem {

        /** ProjectItem id */
        id?: (number|null);

        /** ProjectItem name */
        name?: (string|null);

        /** ProjectItem path */
        path?: (string|null);

        /** ProjectItem web_url */
        web_url?: (string|null);

        /** ProjectItem avatar_url */
        avatar_url?: (string|null);

        /** ProjectItem description */
        description?: (string|null);

        /** ProjectItem enabled */
        enabled?: (boolean|null);

        /** ProjectItem global_enabled */
        global_enabled?: (boolean|null);
    }

    /** Represents a ProjectItem. */
    class ProjectItem implements IProjectItem {

        /**
         * Constructs a new ProjectItem.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IProjectItem);

        /** ProjectItem id. */
        public id: number;

        /** ProjectItem name. */
        public name: string;

        /** ProjectItem path. */
        public path: string;

        /** ProjectItem web_url. */
        public web_url: string;

        /** ProjectItem avatar_url. */
        public avatar_url: string;

        /** ProjectItem description. */
        public description: string;

        /** ProjectItem enabled. */
        public enabled: boolean;

        /** ProjectItem global_enabled. */
        public global_enabled: boolean;

        /**
         * Encodes the specified ProjectItem message. Does not implicitly {@link git.ProjectItem.verify|verify} messages.
         * @param message ProjectItem message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.ProjectItem, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ProjectItem message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ProjectItem
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.ProjectItem;
    }

    /** Properties of an AllProjectsResponse. */
    interface IAllProjectsResponse {

        /** AllProjectsResponse items */
        items?: (git.ProjectItem[]|null);
    }

    /** Represents an AllProjectsResponse. */
    class AllProjectsResponse implements IAllProjectsResponse {

        /**
         * Constructs a new AllProjectsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IAllProjectsResponse);

        /** AllProjectsResponse items. */
        public items: git.ProjectItem[];

        /**
         * Encodes the specified AllProjectsResponse message. Does not implicitly {@link git.AllProjectsResponse.verify|verify} messages.
         * @param message AllProjectsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.AllProjectsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllProjectsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllProjectsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.AllProjectsResponse;
    }

    /** Properties of an Option. */
    interface IOption {

        /** Option value */
        value?: (string|null);

        /** Option label */
        label?: (string|null);

        /** Option type */
        type?: (string|null);

        /** Option isLeaf */
        isLeaf?: (boolean|null);

        /** Option gitProjectId */
        gitProjectId?: (string|null);

        /** Option branch */
        branch?: (string|null);
    }

    /** Represents an Option. */
    class Option implements IOption {

        /**
         * Constructs a new Option.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IOption);

        /** Option value. */
        public value: string;

        /** Option label. */
        public label: string;

        /** Option type. */
        public type: string;

        /** Option isLeaf. */
        public isLeaf: boolean;

        /** Option gitProjectId. */
        public gitProjectId: string;

        /** Option branch. */
        public branch: string;

        /**
         * Encodes the specified Option message. Does not implicitly {@link git.Option.verify|verify} messages.
         * @param message Option message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.Option, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an Option message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Option
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.Option;
    }

    /** Properties of a ProjectOptionsResponse. */
    interface IProjectOptionsResponse {

        /** ProjectOptionsResponse items */
        items?: (git.Option[]|null);
    }

    /** Represents a ProjectOptionsResponse. */
    class ProjectOptionsResponse implements IProjectOptionsResponse {

        /**
         * Constructs a new ProjectOptionsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IProjectOptionsResponse);

        /** ProjectOptionsResponse items. */
        public items: git.Option[];

        /**
         * Encodes the specified ProjectOptionsResponse message. Does not implicitly {@link git.ProjectOptionsResponse.verify|verify} messages.
         * @param message ProjectOptionsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.ProjectOptionsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ProjectOptionsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ProjectOptionsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.ProjectOptionsResponse;
    }

    /** Properties of a BranchOptionsRequest. */
    interface IBranchOptionsRequest {

        /** BranchOptionsRequest git_project_id */
        git_project_id?: (string|null);

        /** BranchOptionsRequest all */
        all?: (boolean|null);
    }

    /** Represents a BranchOptionsRequest. */
    class BranchOptionsRequest implements IBranchOptionsRequest {

        /**
         * Constructs a new BranchOptionsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IBranchOptionsRequest);

        /** BranchOptionsRequest git_project_id. */
        public git_project_id: string;

        /** BranchOptionsRequest all. */
        public all: boolean;

        /**
         * Encodes the specified BranchOptionsRequest message. Does not implicitly {@link git.BranchOptionsRequest.verify|verify} messages.
         * @param message BranchOptionsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.BranchOptionsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BranchOptionsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BranchOptionsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.BranchOptionsRequest;
    }

    /** Properties of a BranchOptionsResponse. */
    interface IBranchOptionsResponse {

        /** BranchOptionsResponse items */
        items?: (git.Option[]|null);
    }

    /** Represents a BranchOptionsResponse. */
    class BranchOptionsResponse implements IBranchOptionsResponse {

        /**
         * Constructs a new BranchOptionsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IBranchOptionsResponse);

        /** BranchOptionsResponse items. */
        public items: git.Option[];

        /**
         * Encodes the specified BranchOptionsResponse message. Does not implicitly {@link git.BranchOptionsResponse.verify|verify} messages.
         * @param message BranchOptionsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.BranchOptionsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BranchOptionsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BranchOptionsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.BranchOptionsResponse;
    }

    /** Properties of a CommitOptionsRequest. */
    interface ICommitOptionsRequest {

        /** CommitOptionsRequest git_project_id */
        git_project_id?: (string|null);

        /** CommitOptionsRequest branch */
        branch?: (string|null);
    }

    /** Represents a CommitOptionsRequest. */
    class CommitOptionsRequest implements ICommitOptionsRequest {

        /**
         * Constructs a new CommitOptionsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.ICommitOptionsRequest);

        /** CommitOptionsRequest git_project_id. */
        public git_project_id: string;

        /** CommitOptionsRequest branch. */
        public branch: string;

        /**
         * Encodes the specified CommitOptionsRequest message. Does not implicitly {@link git.CommitOptionsRequest.verify|verify} messages.
         * @param message CommitOptionsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.CommitOptionsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CommitOptionsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CommitOptionsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.CommitOptionsRequest;
    }

    /** Properties of a CommitOptionsResponse. */
    interface ICommitOptionsResponse {

        /** CommitOptionsResponse items */
        items?: (git.Option[]|null);
    }

    /** Represents a CommitOptionsResponse. */
    class CommitOptionsResponse implements ICommitOptionsResponse {

        /**
         * Constructs a new CommitOptionsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.ICommitOptionsResponse);

        /** CommitOptionsResponse items. */
        public items: git.Option[];

        /**
         * Encodes the specified CommitOptionsResponse message. Does not implicitly {@link git.CommitOptionsResponse.verify|verify} messages.
         * @param message CommitOptionsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.CommitOptionsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CommitOptionsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CommitOptionsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.CommitOptionsResponse;
    }

    /** Properties of a CommitRequest. */
    interface ICommitRequest {

        /** CommitRequest git_project_id */
        git_project_id?: (string|null);

        /** CommitRequest branch */
        branch?: (string|null);

        /** CommitRequest commit */
        commit?: (string|null);
    }

    /** Represents a CommitRequest. */
    class CommitRequest implements ICommitRequest {

        /**
         * Constructs a new CommitRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.ICommitRequest);

        /** CommitRequest git_project_id. */
        public git_project_id: string;

        /** CommitRequest branch. */
        public branch: string;

        /** CommitRequest commit. */
        public commit: string;

        /**
         * Encodes the specified CommitRequest message. Does not implicitly {@link git.CommitRequest.verify|verify} messages.
         * @param message CommitRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.CommitRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CommitRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CommitRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.CommitRequest;
    }

    /** Properties of a CommitResponse. */
    interface ICommitResponse {

        /** CommitResponse id */
        id?: (string|null);

        /** CommitResponse short_id */
        short_id?: (string|null);

        /** CommitResponse git_project_id */
        git_project_id?: (string|null);

        /** CommitResponse label */
        label?: (string|null);

        /** CommitResponse title */
        title?: (string|null);

        /** CommitResponse branch */
        branch?: (string|null);

        /** CommitResponse author_name */
        author_name?: (string|null);

        /** CommitResponse author_email */
        author_email?: (string|null);

        /** CommitResponse committer_name */
        committer_name?: (string|null);

        /** CommitResponse committer_email */
        committer_email?: (string|null);

        /** CommitResponse web_url */
        web_url?: (string|null);

        /** CommitResponse message */
        message?: (string|null);

        /** CommitResponse committed_date */
        committed_date?: (string|null);

        /** CommitResponse created_at */
        created_at?: (string|null);
    }

    /** Represents a CommitResponse. */
    class CommitResponse implements ICommitResponse {

        /**
         * Constructs a new CommitResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.ICommitResponse);

        /** CommitResponse id. */
        public id: string;

        /** CommitResponse short_id. */
        public short_id: string;

        /** CommitResponse git_project_id. */
        public git_project_id: string;

        /** CommitResponse label. */
        public label: string;

        /** CommitResponse title. */
        public title: string;

        /** CommitResponse branch. */
        public branch: string;

        /** CommitResponse author_name. */
        public author_name: string;

        /** CommitResponse author_email. */
        public author_email: string;

        /** CommitResponse committer_name. */
        public committer_name: string;

        /** CommitResponse committer_email. */
        public committer_email: string;

        /** CommitResponse web_url. */
        public web_url: string;

        /** CommitResponse message. */
        public message: string;

        /** CommitResponse committed_date. */
        public committed_date: string;

        /** CommitResponse created_at. */
        public created_at: string;

        /**
         * Encodes the specified CommitResponse message. Does not implicitly {@link git.CommitResponse.verify|verify} messages.
         * @param message CommitResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.CommitResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CommitResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CommitResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.CommitResponse;
    }

    /** Properties of a PipelineInfoRequest. */
    interface IPipelineInfoRequest {

        /** PipelineInfoRequest git_project_id */
        git_project_id?: (string|null);

        /** PipelineInfoRequest branch */
        branch?: (string|null);

        /** PipelineInfoRequest commit */
        commit?: (string|null);
    }

    /** Represents a PipelineInfoRequest. */
    class PipelineInfoRequest implements IPipelineInfoRequest {

        /**
         * Constructs a new PipelineInfoRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IPipelineInfoRequest);

        /** PipelineInfoRequest git_project_id. */
        public git_project_id: string;

        /** PipelineInfoRequest branch. */
        public branch: string;

        /** PipelineInfoRequest commit. */
        public commit: string;

        /**
         * Encodes the specified PipelineInfoRequest message. Does not implicitly {@link git.PipelineInfoRequest.verify|verify} messages.
         * @param message PipelineInfoRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.PipelineInfoRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a PipelineInfoRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns PipelineInfoRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.PipelineInfoRequest;
    }

    /** Properties of a PipelineInfoResponse. */
    interface IPipelineInfoResponse {

        /** PipelineInfoResponse status */
        status?: (string|null);

        /** PipelineInfoResponse web_url */
        web_url?: (string|null);
    }

    /** Represents a PipelineInfoResponse. */
    class PipelineInfoResponse implements IPipelineInfoResponse {

        /**
         * Constructs a new PipelineInfoResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IPipelineInfoResponse);

        /** PipelineInfoResponse status. */
        public status: string;

        /** PipelineInfoResponse web_url. */
        public web_url: string;

        /**
         * Encodes the specified PipelineInfoResponse message. Does not implicitly {@link git.PipelineInfoResponse.verify|verify} messages.
         * @param message PipelineInfoResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.PipelineInfoResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a PipelineInfoResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns PipelineInfoResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.PipelineInfoResponse;
    }

    /** Properties of a ConfigFileRequest. */
    interface IConfigFileRequest {

        /** ConfigFileRequest git_project_id */
        git_project_id?: (string|null);

        /** ConfigFileRequest branch */
        branch?: (string|null);
    }

    /** Represents a ConfigFileRequest. */
    class ConfigFileRequest implements IConfigFileRequest {

        /**
         * Constructs a new ConfigFileRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IConfigFileRequest);

        /** ConfigFileRequest git_project_id. */
        public git_project_id: string;

        /** ConfigFileRequest branch. */
        public branch: string;

        /**
         * Encodes the specified ConfigFileRequest message. Does not implicitly {@link git.ConfigFileRequest.verify|verify} messages.
         * @param message ConfigFileRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.ConfigFileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ConfigFileRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ConfigFileRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.ConfigFileRequest;
    }

    /** Properties of a ConfigFileResponse. */
    interface IConfigFileResponse {

        /** ConfigFileResponse data */
        data?: (string|null);

        /** ConfigFileResponse type */
        type?: (string|null);

        /** ConfigFileResponse elements */
        elements?: (mars.Element[]|null);
    }

    /** Represents a ConfigFileResponse. */
    class ConfigFileResponse implements IConfigFileResponse {

        /**
         * Constructs a new ConfigFileResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IConfigFileResponse);

        /** ConfigFileResponse data. */
        public data: string;

        /** ConfigFileResponse type. */
        public type: string;

        /** ConfigFileResponse elements. */
        public elements: mars.Element[];

        /**
         * Encodes the specified ConfigFileResponse message. Does not implicitly {@link git.ConfigFileResponse.verify|verify} messages.
         * @param message ConfigFileResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.ConfigFileResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ConfigFileResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ConfigFileResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.ConfigFileResponse;
    }

    /** Properties of an EnableProjectResponse. */
    interface IEnableProjectResponse {
    }

    /** Represents an EnableProjectResponse. */
    class EnableProjectResponse implements IEnableProjectResponse {

        /**
         * Constructs a new EnableProjectResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IEnableProjectResponse);

        /**
         * Encodes the specified EnableProjectResponse message. Does not implicitly {@link git.EnableProjectResponse.verify|verify} messages.
         * @param message EnableProjectResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.EnableProjectResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EnableProjectResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns EnableProjectResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.EnableProjectResponse;
    }

    /** Properties of a DisableProjectResponse. */
    interface IDisableProjectResponse {
    }

    /** Represents a DisableProjectResponse. */
    class DisableProjectResponse implements IDisableProjectResponse {

        /**
         * Constructs a new DisableProjectResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IDisableProjectResponse);

        /**
         * Encodes the specified DisableProjectResponse message. Does not implicitly {@link git.DisableProjectResponse.verify|verify} messages.
         * @param message DisableProjectResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.DisableProjectResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DisableProjectResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DisableProjectResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.DisableProjectResponse;
    }

    /** Properties of an AllProjectsRequest. */
    interface IAllProjectsRequest {
    }

    /** Represents an AllProjectsRequest. */
    class AllProjectsRequest implements IAllProjectsRequest {

        /**
         * Constructs a new AllProjectsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IAllProjectsRequest);

        /**
         * Encodes the specified AllProjectsRequest message. Does not implicitly {@link git.AllProjectsRequest.verify|verify} messages.
         * @param message AllProjectsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.AllProjectsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllProjectsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllProjectsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.AllProjectsRequest;
    }

    /** Properties of a ProjectOptionsRequest. */
    interface IProjectOptionsRequest {
    }

    /** Represents a ProjectOptionsRequest. */
    class ProjectOptionsRequest implements IProjectOptionsRequest {

        /**
         * Constructs a new ProjectOptionsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: git.IProjectOptionsRequest);

        /**
         * Encodes the specified ProjectOptionsRequest message. Does not implicitly {@link git.ProjectOptionsRequest.verify|verify} messages.
         * @param message ProjectOptionsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: git.ProjectOptionsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ProjectOptionsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ProjectOptionsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): git.ProjectOptionsRequest;
    }

    /** Represents a Git */
    class Git extends $protobuf.rpc.Service {

        /**
         * Constructs a new Git service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls EnableProject.
         * @param request EnableProjectRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and EnableProjectResponse
         */
        public enableProject(request: git.EnableProjectRequest, callback: git.Git.EnableProjectCallback): void;

        /**
         * Calls EnableProject.
         * @param request EnableProjectRequest message or plain object
         * @returns Promise
         */
        public enableProject(request: git.EnableProjectRequest): Promise<git.EnableProjectResponse>;

        /**
         * Calls DisableProject.
         * @param request DisableProjectRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DisableProjectResponse
         */
        public disableProject(request: git.DisableProjectRequest, callback: git.Git.DisableProjectCallback): void;

        /**
         * Calls DisableProject.
         * @param request DisableProjectRequest message or plain object
         * @returns Promise
         */
        public disableProject(request: git.DisableProjectRequest): Promise<git.DisableProjectResponse>;

        /**
         * Calls All.
         * @param request AllProjectsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and AllProjectsResponse
         */
        public all(request: git.AllProjectsRequest, callback: git.Git.AllCallback): void;

        /**
         * Calls All.
         * @param request AllProjectsRequest message or plain object
         * @returns Promise
         */
        public all(request: git.AllProjectsRequest): Promise<git.AllProjectsResponse>;

        /**
         * Calls ProjectOptions.
         * @param request ProjectOptionsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ProjectOptionsResponse
         */
        public projectOptions(request: git.ProjectOptionsRequest, callback: git.Git.ProjectOptionsCallback): void;

        /**
         * Calls ProjectOptions.
         * @param request ProjectOptionsRequest message or plain object
         * @returns Promise
         */
        public projectOptions(request: git.ProjectOptionsRequest): Promise<git.ProjectOptionsResponse>;

        /**
         * Calls BranchOptions.
         * @param request BranchOptionsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and BranchOptionsResponse
         */
        public branchOptions(request: git.BranchOptionsRequest, callback: git.Git.BranchOptionsCallback): void;

        /**
         * Calls BranchOptions.
         * @param request BranchOptionsRequest message or plain object
         * @returns Promise
         */
        public branchOptions(request: git.BranchOptionsRequest): Promise<git.BranchOptionsResponse>;

        /**
         * Calls CommitOptions.
         * @param request CommitOptionsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CommitOptionsResponse
         */
        public commitOptions(request: git.CommitOptionsRequest, callback: git.Git.CommitOptionsCallback): void;

        /**
         * Calls CommitOptions.
         * @param request CommitOptionsRequest message or plain object
         * @returns Promise
         */
        public commitOptions(request: git.CommitOptionsRequest): Promise<git.CommitOptionsResponse>;

        /**
         * Calls Commit.
         * @param request CommitRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CommitResponse
         */
        public commit(request: git.CommitRequest, callback: git.Git.CommitCallback): void;

        /**
         * Calls Commit.
         * @param request CommitRequest message or plain object
         * @returns Promise
         */
        public commit(request: git.CommitRequest): Promise<git.CommitResponse>;

        /**
         * Calls PipelineInfo.
         * @param request PipelineInfoRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and PipelineInfoResponse
         */
        public pipelineInfo(request: git.PipelineInfoRequest, callback: git.Git.PipelineInfoCallback): void;

        /**
         * Calls PipelineInfo.
         * @param request PipelineInfoRequest message or plain object
         * @returns Promise
         */
        public pipelineInfo(request: git.PipelineInfoRequest): Promise<git.PipelineInfoResponse>;

        /**
         * Calls MarsConfigFile.
         * @param request ConfigFileRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ConfigFileResponse
         */
        public marsConfigFile(request: git.ConfigFileRequest, callback: git.Git.MarsConfigFileCallback): void;

        /**
         * Calls MarsConfigFile.
         * @param request ConfigFileRequest message or plain object
         * @returns Promise
         */
        public marsConfigFile(request: git.ConfigFileRequest): Promise<git.ConfigFileResponse>;
    }

    namespace Git {

        /**
         * Callback as used by {@link git.Git#enableProject}.
         * @param error Error, if any
         * @param [response] EnableProjectResponse
         */
        type EnableProjectCallback = (error: (Error|null), response?: git.EnableProjectResponse) => void;

        /**
         * Callback as used by {@link git.Git#disableProject}.
         * @param error Error, if any
         * @param [response] DisableProjectResponse
         */
        type DisableProjectCallback = (error: (Error|null), response?: git.DisableProjectResponse) => void;

        /**
         * Callback as used by {@link git.Git#all}.
         * @param error Error, if any
         * @param [response] AllProjectsResponse
         */
        type AllCallback = (error: (Error|null), response?: git.AllProjectsResponse) => void;

        /**
         * Callback as used by {@link git.Git#projectOptions}.
         * @param error Error, if any
         * @param [response] ProjectOptionsResponse
         */
        type ProjectOptionsCallback = (error: (Error|null), response?: git.ProjectOptionsResponse) => void;

        /**
         * Callback as used by {@link git.Git#branchOptions}.
         * @param error Error, if any
         * @param [response] BranchOptionsResponse
         */
        type BranchOptionsCallback = (error: (Error|null), response?: git.BranchOptionsResponse) => void;

        /**
         * Callback as used by {@link git.Git#commitOptions}.
         * @param error Error, if any
         * @param [response] CommitOptionsResponse
         */
        type CommitOptionsCallback = (error: (Error|null), response?: git.CommitOptionsResponse) => void;

        /**
         * Callback as used by {@link git.Git#commit}.
         * @param error Error, if any
         * @param [response] CommitResponse
         */
        type CommitCallback = (error: (Error|null), response?: git.CommitResponse) => void;

        /**
         * Callback as used by {@link git.Git#pipelineInfo}.
         * @param error Error, if any
         * @param [response] PipelineInfoResponse
         */
        type PipelineInfoCallback = (error: (Error|null), response?: git.PipelineInfoResponse) => void;

        /**
         * Callback as used by {@link git.Git#marsConfigFile}.
         * @param error Error, if any
         * @param [response] ConfigFileResponse
         */
        type MarsConfigFileCallback = (error: (Error|null), response?: git.ConfigFileResponse) => void;
    }
}

/** Namespace gitconfig. */
export namespace gitconfig {

    /** Properties of a FileRequest. */
    interface IFileRequest {

        /** FileRequest git_project_id */
        git_project_id?: (string|null);

        /** FileRequest branch */
        branch?: (string|null);
    }

    /** Represents a FileRequest. */
    class FileRequest implements IFileRequest {

        /**
         * Constructs a new FileRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IFileRequest);

        /** FileRequest git_project_id. */
        public git_project_id: string;

        /** FileRequest branch. */
        public branch: string;

        /**
         * Encodes the specified FileRequest message. Does not implicitly {@link gitconfig.FileRequest.verify|verify} messages.
         * @param message FileRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.FileRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a FileRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns FileRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.FileRequest;
    }

    /** Properties of a FileResponse. */
    interface IFileResponse {

        /** FileResponse data */
        data?: (string|null);

        /** FileResponse type */
        type?: (string|null);

        /** FileResponse elements */
        elements?: (mars.Element[]|null);
    }

    /** Represents a FileResponse. */
    class FileResponse implements IFileResponse {

        /**
         * Constructs a new FileResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IFileResponse);

        /** FileResponse data. */
        public data: string;

        /** FileResponse type. */
        public type: string;

        /** FileResponse elements. */
        public elements: mars.Element[];

        /**
         * Encodes the specified FileResponse message. Does not implicitly {@link gitconfig.FileResponse.verify|verify} messages.
         * @param message FileResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.FileResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a FileResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns FileResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.FileResponse;
    }

    /** Properties of a ShowRequest. */
    interface IShowRequest {

        /** ShowRequest git_project_id */
        git_project_id?: (number|null);

        /** ShowRequest branch */
        branch?: (string|null);
    }

    /** Represents a ShowRequest. */
    class ShowRequest implements IShowRequest {

        /**
         * Constructs a new ShowRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IShowRequest);

        /** ShowRequest git_project_id. */
        public git_project_id: number;

        /** ShowRequest branch. */
        public branch: string;

        /**
         * Encodes the specified ShowRequest message. Does not implicitly {@link gitconfig.ShowRequest.verify|verify} messages.
         * @param message ShowRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.ShowRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.ShowRequest;
    }

    /** Properties of a ShowResponse. */
    interface IShowResponse {

        /** ShowResponse branch */
        branch?: (string|null);

        /** ShowResponse config */
        config?: (mars.Config|null);
    }

    /** Represents a ShowResponse. */
    class ShowResponse implements IShowResponse {

        /**
         * Constructs a new ShowResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IShowResponse);

        /** ShowResponse branch. */
        public branch: string;

        /** ShowResponse config. */
        public config?: (mars.Config|null);

        /**
         * Encodes the specified ShowResponse message. Does not implicitly {@link gitconfig.ShowResponse.verify|verify} messages.
         * @param message ShowResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.ShowResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.ShowResponse;
    }

    /** Properties of a GlobalConfigRequest. */
    interface IGlobalConfigRequest {

        /** GlobalConfigRequest git_project_id */
        git_project_id?: (number|null);
    }

    /** Represents a GlobalConfigRequest. */
    class GlobalConfigRequest implements IGlobalConfigRequest {

        /**
         * Constructs a new GlobalConfigRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IGlobalConfigRequest);

        /** GlobalConfigRequest git_project_id. */
        public git_project_id: number;

        /**
         * Encodes the specified GlobalConfigRequest message. Does not implicitly {@link gitconfig.GlobalConfigRequest.verify|verify} messages.
         * @param message GlobalConfigRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.GlobalConfigRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GlobalConfigRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GlobalConfigRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.GlobalConfigRequest;
    }

    /** Properties of a GlobalConfigResponse. */
    interface IGlobalConfigResponse {

        /** GlobalConfigResponse enabled */
        enabled?: (boolean|null);

        /** GlobalConfigResponse config */
        config?: (mars.Config|null);
    }

    /** Represents a GlobalConfigResponse. */
    class GlobalConfigResponse implements IGlobalConfigResponse {

        /**
         * Constructs a new GlobalConfigResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IGlobalConfigResponse);

        /** GlobalConfigResponse enabled. */
        public enabled: boolean;

        /** GlobalConfigResponse config. */
        public config?: (mars.Config|null);

        /**
         * Encodes the specified GlobalConfigResponse message. Does not implicitly {@link gitconfig.GlobalConfigResponse.verify|verify} messages.
         * @param message GlobalConfigResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.GlobalConfigResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GlobalConfigResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GlobalConfigResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.GlobalConfigResponse;
    }

    /** Properties of an UpdateRequest. */
    interface IUpdateRequest {

        /** UpdateRequest git_project_id */
        git_project_id?: (number|null);

        /** UpdateRequest config */
        config?: (mars.Config|null);
    }

    /** Represents an UpdateRequest. */
    class UpdateRequest implements IUpdateRequest {

        /**
         * Constructs a new UpdateRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IUpdateRequest);

        /** UpdateRequest git_project_id. */
        public git_project_id: number;

        /** UpdateRequest config. */
        public config?: (mars.Config|null);

        /**
         * Encodes the specified UpdateRequest message. Does not implicitly {@link gitconfig.UpdateRequest.verify|verify} messages.
         * @param message UpdateRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.UpdateRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an UpdateRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns UpdateRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.UpdateRequest;
    }

    /** Properties of an UpdateResponse. */
    interface IUpdateResponse {

        /** UpdateResponse config */
        config?: (mars.Config|null);
    }

    /** Represents an UpdateResponse. */
    class UpdateResponse implements IUpdateResponse {

        /**
         * Constructs a new UpdateResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IUpdateResponse);

        /** UpdateResponse config. */
        public config?: (mars.Config|null);

        /**
         * Encodes the specified UpdateResponse message. Does not implicitly {@link gitconfig.UpdateResponse.verify|verify} messages.
         * @param message UpdateResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.UpdateResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an UpdateResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns UpdateResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.UpdateResponse;
    }

    /** Properties of a ToggleGlobalStatusRequest. */
    interface IToggleGlobalStatusRequest {

        /** ToggleGlobalStatusRequest git_project_id */
        git_project_id?: (number|null);

        /** ToggleGlobalStatusRequest enabled */
        enabled?: (boolean|null);
    }

    /** Represents a ToggleGlobalStatusRequest. */
    class ToggleGlobalStatusRequest implements IToggleGlobalStatusRequest {

        /**
         * Constructs a new ToggleGlobalStatusRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IToggleGlobalStatusRequest);

        /** ToggleGlobalStatusRequest git_project_id. */
        public git_project_id: number;

        /** ToggleGlobalStatusRequest enabled. */
        public enabled: boolean;

        /**
         * Encodes the specified ToggleGlobalStatusRequest message. Does not implicitly {@link gitconfig.ToggleGlobalStatusRequest.verify|verify} messages.
         * @param message ToggleGlobalStatusRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.ToggleGlobalStatusRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ToggleGlobalStatusRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ToggleGlobalStatusRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.ToggleGlobalStatusRequest;
    }

    /** Properties of a DefaultChartValuesRequest. */
    interface IDefaultChartValuesRequest {

        /** DefaultChartValuesRequest git_project_id */
        git_project_id?: (number|null);

        /** DefaultChartValuesRequest branch */
        branch?: (string|null);
    }

    /** Represents a DefaultChartValuesRequest. */
    class DefaultChartValuesRequest implements IDefaultChartValuesRequest {

        /**
         * Constructs a new DefaultChartValuesRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IDefaultChartValuesRequest);

        /** DefaultChartValuesRequest git_project_id. */
        public git_project_id: number;

        /** DefaultChartValuesRequest branch. */
        public branch: string;

        /**
         * Encodes the specified DefaultChartValuesRequest message. Does not implicitly {@link gitconfig.DefaultChartValuesRequest.verify|verify} messages.
         * @param message DefaultChartValuesRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.DefaultChartValuesRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DefaultChartValuesRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DefaultChartValuesRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.DefaultChartValuesRequest;
    }

    /** Properties of a DefaultChartValuesResponse. */
    interface IDefaultChartValuesResponse {

        /** DefaultChartValuesResponse value */
        value?: (string|null);
    }

    /** Represents a DefaultChartValuesResponse. */
    class DefaultChartValuesResponse implements IDefaultChartValuesResponse {

        /**
         * Constructs a new DefaultChartValuesResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IDefaultChartValuesResponse);

        /** DefaultChartValuesResponse value. */
        public value: string;

        /**
         * Encodes the specified DefaultChartValuesResponse message. Does not implicitly {@link gitconfig.DefaultChartValuesResponse.verify|verify} messages.
         * @param message DefaultChartValuesResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.DefaultChartValuesResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DefaultChartValuesResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DefaultChartValuesResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.DefaultChartValuesResponse;
    }

    /** Properties of a ToggleGlobalStatusResponse. */
    interface IToggleGlobalStatusResponse {
    }

    /** Represents a ToggleGlobalStatusResponse. */
    class ToggleGlobalStatusResponse implements IToggleGlobalStatusResponse {

        /**
         * Constructs a new ToggleGlobalStatusResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: gitconfig.IToggleGlobalStatusResponse);

        /**
         * Encodes the specified ToggleGlobalStatusResponse message. Does not implicitly {@link gitconfig.ToggleGlobalStatusResponse.verify|verify} messages.
         * @param message ToggleGlobalStatusResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: gitconfig.ToggleGlobalStatusResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ToggleGlobalStatusResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ToggleGlobalStatusResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): gitconfig.ToggleGlobalStatusResponse;
    }

    /** Represents a GitConfig */
    class GitConfig extends $protobuf.rpc.Service {

        /**
         * Constructs a new GitConfig service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ShowResponse
         */
        public show(request: gitconfig.ShowRequest, callback: gitconfig.GitConfig.ShowCallback): void;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @returns Promise
         */
        public show(request: gitconfig.ShowRequest): Promise<gitconfig.ShowResponse>;

        /**
         * Calls GlobalConfig.
         * @param request GlobalConfigRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and GlobalConfigResponse
         */
        public globalConfig(request: gitconfig.GlobalConfigRequest, callback: gitconfig.GitConfig.GlobalConfigCallback): void;

        /**
         * Calls GlobalConfig.
         * @param request GlobalConfigRequest message or plain object
         * @returns Promise
         */
        public globalConfig(request: gitconfig.GlobalConfigRequest): Promise<gitconfig.GlobalConfigResponse>;

        /**
         * Calls ToggleGlobalStatus.
         * @param request ToggleGlobalStatusRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ToggleGlobalStatusResponse
         */
        public toggleGlobalStatus(request: gitconfig.ToggleGlobalStatusRequest, callback: gitconfig.GitConfig.ToggleGlobalStatusCallback): void;

        /**
         * Calls ToggleGlobalStatus.
         * @param request ToggleGlobalStatusRequest message or plain object
         * @returns Promise
         */
        public toggleGlobalStatus(request: gitconfig.ToggleGlobalStatusRequest): Promise<gitconfig.ToggleGlobalStatusResponse>;

        /**
         * Calls Update.
         * @param request UpdateRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and UpdateResponse
         */
        public update(request: gitconfig.UpdateRequest, callback: gitconfig.GitConfig.UpdateCallback): void;

        /**
         * Calls Update.
         * @param request UpdateRequest message or plain object
         * @returns Promise
         */
        public update(request: gitconfig.UpdateRequest): Promise<gitconfig.UpdateResponse>;

        /**
         * Calls GetDefaultChartValues.
         * @param request DefaultChartValuesRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DefaultChartValuesResponse
         */
        public getDefaultChartValues(request: gitconfig.DefaultChartValuesRequest, callback: gitconfig.GitConfig.GetDefaultChartValuesCallback): void;

        /**
         * Calls GetDefaultChartValues.
         * @param request DefaultChartValuesRequest message or plain object
         * @returns Promise
         */
        public getDefaultChartValues(request: gitconfig.DefaultChartValuesRequest): Promise<gitconfig.DefaultChartValuesResponse>;
    }

    namespace GitConfig {

        /**
         * Callback as used by {@link gitconfig.GitConfig#show}.
         * @param error Error, if any
         * @param [response] ShowResponse
         */
        type ShowCallback = (error: (Error|null), response?: gitconfig.ShowResponse) => void;

        /**
         * Callback as used by {@link gitconfig.GitConfig#globalConfig}.
         * @param error Error, if any
         * @param [response] GlobalConfigResponse
         */
        type GlobalConfigCallback = (error: (Error|null), response?: gitconfig.GlobalConfigResponse) => void;

        /**
         * Callback as used by {@link gitconfig.GitConfig#toggleGlobalStatus}.
         * @param error Error, if any
         * @param [response] ToggleGlobalStatusResponse
         */
        type ToggleGlobalStatusCallback = (error: (Error|null), response?: gitconfig.ToggleGlobalStatusResponse) => void;

        /**
         * Callback as used by {@link gitconfig.GitConfig#update}.
         * @param error Error, if any
         * @param [response] UpdateResponse
         */
        type UpdateCallback = (error: (Error|null), response?: gitconfig.UpdateResponse) => void;

        /**
         * Callback as used by {@link gitconfig.GitConfig#getDefaultChartValues}.
         * @param error Error, if any
         * @param [response] DefaultChartValuesResponse
         */
        type GetDefaultChartValuesCallback = (error: (Error|null), response?: gitconfig.DefaultChartValuesResponse) => void;
    }
}

/** Namespace mars. */
export namespace mars {

    /** Properties of a Config. */
    interface IConfig {

        /** Config config_file */
        config_file?: (string|null);

        /** Config config_file_values */
        config_file_values?: (string|null);

        /** Config config_field */
        config_field?: (string|null);

        /** Config is_simple_env */
        is_simple_env?: (boolean|null);

        /** Config config_file_type */
        config_file_type?: (string|null);

        /** Config local_chart_path */
        local_chart_path?: (string|null);

        /** Config branches */
        branches?: (string[]|null);

        /** Config values_yaml */
        values_yaml?: (string|null);

        /** Config elements */
        elements?: (mars.Element[]|null);
    }

    /** Represents a Config. */
    class Config implements IConfig {

        /**
         * Constructs a new Config.
         * @param [properties] Properties to set
         */
        constructor(properties?: mars.IConfig);

        /** Config config_file. */
        public config_file: string;

        /** Config config_file_values. */
        public config_file_values: string;

        /** Config config_field. */
        public config_field: string;

        /** Config is_simple_env. */
        public is_simple_env: boolean;

        /** Config config_file_type. */
        public config_file_type: string;

        /** Config local_chart_path. */
        public local_chart_path: string;

        /** Config branches. */
        public branches: string[];

        /** Config values_yaml. */
        public values_yaml: string;

        /** Config elements. */
        public elements: mars.Element[];

        /**
         * Encodes the specified Config message. Does not implicitly {@link mars.Config.verify|verify} messages.
         * @param message Config message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: mars.Config, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Config message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Config
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): mars.Config;
    }

    /** ElementType enum. */
    enum ElementType {
        ElementTypeUnknown = 0,
        ElementTypeInput = 1,
        ElementTypeInputNumber = 2,
        ElementTypeSelect = 3,
        ElementTypeRadio = 4,
        ElementTypeSwitch = 5
    }

    /** Properties of an Element. */
    interface IElement {

        /** Element path */
        path?: (string|null);

        /** Element type */
        type?: (mars.ElementType|null);

        /** Element default */
        "default"?: (string|null);

        /** Element description */
        description?: (string|null);

        /** Element select_values */
        select_values?: (string[]|null);
    }

    /** Represents an Element. */
    class Element implements IElement {

        /**
         * Constructs a new Element.
         * @param [properties] Properties to set
         */
        constructor(properties?: mars.IElement);

        /** Element path. */
        public path: string;

        /** Element type. */
        public type: mars.ElementType;

        /** Element default. */
        public default: string;

        /** Element description. */
        public description: string;

        /** Element select_values. */
        public select_values: string[];

        /**
         * Encodes the specified Element message. Does not implicitly {@link mars.Element.verify|verify} messages.
         * @param message Element message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: mars.Element, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an Element message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Element
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): mars.Element;
    }
}

/** Namespace metrics. */
export namespace metrics {

    /** Properties of a TopPodRequest. */
    interface ITopPodRequest {

        /** TopPodRequest namespace */
        namespace?: (string|null);

        /** TopPodRequest pod */
        pod?: (string|null);
    }

    /** Represents a TopPodRequest. */
    class TopPodRequest implements ITopPodRequest {

        /**
         * Constructs a new TopPodRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ITopPodRequest);

        /** TopPodRequest namespace. */
        public namespace: string;

        /** TopPodRequest pod. */
        public pod: string;

        /**
         * Encodes the specified TopPodRequest message. Does not implicitly {@link metrics.TopPodRequest.verify|verify} messages.
         * @param message TopPodRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.TopPodRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TopPodRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TopPodRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.TopPodRequest;
    }

    /** Properties of a TopPodResponse. */
    interface ITopPodResponse {

        /** TopPodResponse cpu */
        cpu?: (number|null);

        /** TopPodResponse memory */
        memory?: (number|null);

        /** TopPodResponse humanize_cpu */
        humanize_cpu?: (string|null);

        /** TopPodResponse humanize_memory */
        humanize_memory?: (string|null);

        /** TopPodResponse time */
        time?: (string|null);

        /** TopPodResponse length */
        length?: (number|null);
    }

    /** Represents a TopPodResponse. */
    class TopPodResponse implements ITopPodResponse {

        /**
         * Constructs a new TopPodResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ITopPodResponse);

        /** TopPodResponse cpu. */
        public cpu: number;

        /** TopPodResponse memory. */
        public memory: number;

        /** TopPodResponse humanize_cpu. */
        public humanize_cpu: string;

        /** TopPodResponse humanize_memory. */
        public humanize_memory: string;

        /** TopPodResponse time. */
        public time: string;

        /** TopPodResponse length. */
        public length: number;

        /**
         * Encodes the specified TopPodResponse message. Does not implicitly {@link metrics.TopPodResponse.verify|verify} messages.
         * @param message TopPodResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.TopPodResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TopPodResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TopPodResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.TopPodResponse;
    }

    /** Properties of a CpuMemoryInNamespaceRequest. */
    interface ICpuMemoryInNamespaceRequest {

        /** CpuMemoryInNamespaceRequest namespace_id */
        namespace_id?: (number|null);
    }

    /** Represents a CpuMemoryInNamespaceRequest. */
    class CpuMemoryInNamespaceRequest implements ICpuMemoryInNamespaceRequest {

        /**
         * Constructs a new CpuMemoryInNamespaceRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ICpuMemoryInNamespaceRequest);

        /** CpuMemoryInNamespaceRequest namespace_id. */
        public namespace_id: number;

        /**
         * Encodes the specified CpuMemoryInNamespaceRequest message. Does not implicitly {@link metrics.CpuMemoryInNamespaceRequest.verify|verify} messages.
         * @param message CpuMemoryInNamespaceRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.CpuMemoryInNamespaceRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CpuMemoryInNamespaceRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CpuMemoryInNamespaceRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.CpuMemoryInNamespaceRequest;
    }

    /** Properties of a CpuMemoryInNamespaceResponse. */
    interface ICpuMemoryInNamespaceResponse {

        /** CpuMemoryInNamespaceResponse cpu */
        cpu?: (string|null);

        /** CpuMemoryInNamespaceResponse memory */
        memory?: (string|null);
    }

    /** Represents a CpuMemoryInNamespaceResponse. */
    class CpuMemoryInNamespaceResponse implements ICpuMemoryInNamespaceResponse {

        /**
         * Constructs a new CpuMemoryInNamespaceResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ICpuMemoryInNamespaceResponse);

        /** CpuMemoryInNamespaceResponse cpu. */
        public cpu: string;

        /** CpuMemoryInNamespaceResponse memory. */
        public memory: string;

        /**
         * Encodes the specified CpuMemoryInNamespaceResponse message. Does not implicitly {@link metrics.CpuMemoryInNamespaceResponse.verify|verify} messages.
         * @param message CpuMemoryInNamespaceResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.CpuMemoryInNamespaceResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CpuMemoryInNamespaceResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CpuMemoryInNamespaceResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.CpuMemoryInNamespaceResponse;
    }

    /** Properties of a CpuMemoryInProjectRequest. */
    interface ICpuMemoryInProjectRequest {

        /** CpuMemoryInProjectRequest project_id */
        project_id?: (number|null);
    }

    /** Represents a CpuMemoryInProjectRequest. */
    class CpuMemoryInProjectRequest implements ICpuMemoryInProjectRequest {

        /**
         * Constructs a new CpuMemoryInProjectRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ICpuMemoryInProjectRequest);

        /** CpuMemoryInProjectRequest project_id. */
        public project_id: number;

        /**
         * Encodes the specified CpuMemoryInProjectRequest message. Does not implicitly {@link metrics.CpuMemoryInProjectRequest.verify|verify} messages.
         * @param message CpuMemoryInProjectRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.CpuMemoryInProjectRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CpuMemoryInProjectRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CpuMemoryInProjectRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.CpuMemoryInProjectRequest;
    }

    /** Properties of a CpuMemoryInProjectResponse. */
    interface ICpuMemoryInProjectResponse {

        /** CpuMemoryInProjectResponse cpu */
        cpu?: (string|null);

        /** CpuMemoryInProjectResponse memory */
        memory?: (string|null);
    }

    /** Represents a CpuMemoryInProjectResponse. */
    class CpuMemoryInProjectResponse implements ICpuMemoryInProjectResponse {

        /**
         * Constructs a new CpuMemoryInProjectResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: metrics.ICpuMemoryInProjectResponse);

        /** CpuMemoryInProjectResponse cpu. */
        public cpu: string;

        /** CpuMemoryInProjectResponse memory. */
        public memory: string;

        /**
         * Encodes the specified CpuMemoryInProjectResponse message. Does not implicitly {@link metrics.CpuMemoryInProjectResponse.verify|verify} messages.
         * @param message CpuMemoryInProjectResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: metrics.CpuMemoryInProjectResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CpuMemoryInProjectResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CpuMemoryInProjectResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): metrics.CpuMemoryInProjectResponse;
    }

    /** Represents a Metrics */
    class Metrics extends $protobuf.rpc.Service {

        /**
         * Constructs a new Metrics service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls CpuMemoryInNamespace.
         * @param request CpuMemoryInNamespaceRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CpuMemoryInNamespaceResponse
         */
        public cpuMemoryInNamespace(request: metrics.CpuMemoryInNamespaceRequest, callback: metrics.Metrics.CpuMemoryInNamespaceCallback): void;

        /**
         * Calls CpuMemoryInNamespace.
         * @param request CpuMemoryInNamespaceRequest message or plain object
         * @returns Promise
         */
        public cpuMemoryInNamespace(request: metrics.CpuMemoryInNamespaceRequest): Promise<metrics.CpuMemoryInNamespaceResponse>;

        /**
         * Calls CpuMemoryInProject.
         * @param request CpuMemoryInProjectRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CpuMemoryInProjectResponse
         */
        public cpuMemoryInProject(request: metrics.CpuMemoryInProjectRequest, callback: metrics.Metrics.CpuMemoryInProjectCallback): void;

        /**
         * Calls CpuMemoryInProject.
         * @param request CpuMemoryInProjectRequest message or plain object
         * @returns Promise
         */
        public cpuMemoryInProject(request: metrics.CpuMemoryInProjectRequest): Promise<metrics.CpuMemoryInProjectResponse>;

        /**
         * Calls TopPod.
         * @param request TopPodRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and TopPodResponse
         */
        public topPod(request: metrics.TopPodRequest, callback: metrics.Metrics.TopPodCallback): void;

        /**
         * Calls TopPod.
         * @param request TopPodRequest message or plain object
         * @returns Promise
         */
        public topPod(request: metrics.TopPodRequest): Promise<metrics.TopPodResponse>;

        /**
         * Calls StreamTopPod.
         * @param request TopPodRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and TopPodResponse
         */
        public streamTopPod(request: metrics.TopPodRequest, callback: metrics.Metrics.StreamTopPodCallback): void;

        /**
         * Calls StreamTopPod.
         * @param request TopPodRequest message or plain object
         * @returns Promise
         */
        public streamTopPod(request: metrics.TopPodRequest): Promise<metrics.TopPodResponse>;
    }

    namespace Metrics {

        /**
         * Callback as used by {@link metrics.Metrics#cpuMemoryInNamespace}.
         * @param error Error, if any
         * @param [response] CpuMemoryInNamespaceResponse
         */
        type CpuMemoryInNamespaceCallback = (error: (Error|null), response?: metrics.CpuMemoryInNamespaceResponse) => void;

        /**
         * Callback as used by {@link metrics.Metrics#cpuMemoryInProject}.
         * @param error Error, if any
         * @param [response] CpuMemoryInProjectResponse
         */
        type CpuMemoryInProjectCallback = (error: (Error|null), response?: metrics.CpuMemoryInProjectResponse) => void;

        /**
         * Callback as used by {@link metrics.Metrics#topPod}.
         * @param error Error, if any
         * @param [response] TopPodResponse
         */
        type TopPodCallback = (error: (Error|null), response?: metrics.TopPodResponse) => void;

        /**
         * Callback as used by {@link metrics.Metrics#streamTopPod}.
         * @param error Error, if any
         * @param [response] TopPodResponse
         */
        type StreamTopPodCallback = (error: (Error|null), response?: metrics.TopPodResponse) => void;
    }
}

/** Namespace namespace. */
export namespace namespace {

    /** Properties of a CreateRequest. */
    interface ICreateRequest {

        /** CreateRequest namespace */
        namespace?: (string|null);

        /** CreateRequest ignore_if_exists */
        ignore_if_exists?: (boolean|null);
    }

    /** Represents a CreateRequest. */
    class CreateRequest implements ICreateRequest {

        /**
         * Constructs a new CreateRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.ICreateRequest);

        /** CreateRequest namespace. */
        public namespace: string;

        /** CreateRequest ignore_if_exists. */
        public ignore_if_exists: boolean;

        /**
         * Encodes the specified CreateRequest message. Does not implicitly {@link namespace.CreateRequest.verify|verify} messages.
         * @param message CreateRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.CreateRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CreateRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.CreateRequest;
    }

    /** Properties of a ShowRequest. */
    interface IShowRequest {

        /** ShowRequest namespace_id */
        namespace_id?: (number|null);
    }

    /** Represents a ShowRequest. */
    class ShowRequest implements IShowRequest {

        /**
         * Constructs a new ShowRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IShowRequest);

        /** ShowRequest namespace_id. */
        public namespace_id: number;

        /**
         * Encodes the specified ShowRequest message. Does not implicitly {@link namespace.ShowRequest.verify|verify} messages.
         * @param message ShowRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.ShowRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.ShowRequest;
    }

    /** Properties of a DeleteRequest. */
    interface IDeleteRequest {

        /** DeleteRequest namespace_id */
        namespace_id?: (number|null);
    }

    /** Represents a DeleteRequest. */
    class DeleteRequest implements IDeleteRequest {

        /**
         * Constructs a new DeleteRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IDeleteRequest);

        /** DeleteRequest namespace_id. */
        public namespace_id: number;

        /**
         * Encodes the specified DeleteRequest message. Does not implicitly {@link namespace.DeleteRequest.verify|verify} messages.
         * @param message DeleteRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.DeleteRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.DeleteRequest;
    }

    /** Properties of an IsExistsRequest. */
    interface IIsExistsRequest {

        /** IsExistsRequest name */
        name?: (string|null);
    }

    /** Represents an IsExistsRequest. */
    class IsExistsRequest implements IIsExistsRequest {

        /**
         * Constructs a new IsExistsRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IIsExistsRequest);

        /** IsExistsRequest name. */
        public name: string;

        /**
         * Encodes the specified IsExistsRequest message. Does not implicitly {@link namespace.IsExistsRequest.verify|verify} messages.
         * @param message IsExistsRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.IsExistsRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsExistsRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsExistsRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.IsExistsRequest;
    }

    /** Properties of an AllResponse. */
    interface IAllResponse {

        /** AllResponse items */
        items?: (types.NamespaceModel[]|null);
    }

    /** Represents an AllResponse. */
    class AllResponse implements IAllResponse {

        /**
         * Constructs a new AllResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IAllResponse);

        /** AllResponse items. */
        public items: types.NamespaceModel[];

        /**
         * Encodes the specified AllResponse message. Does not implicitly {@link namespace.AllResponse.verify|verify} messages.
         * @param message AllResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.AllResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.AllResponse;
    }

    /** Properties of a CreateResponse. */
    interface ICreateResponse {

        /** CreateResponse namespace */
        namespace?: (types.NamespaceModel|null);

        /** CreateResponse exists */
        exists?: (boolean|null);
    }

    /** Represents a CreateResponse. */
    class CreateResponse implements ICreateResponse {

        /**
         * Constructs a new CreateResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.ICreateResponse);

        /** CreateResponse namespace. */
        public namespace?: (types.NamespaceModel|null);

        /** CreateResponse exists. */
        public exists: boolean;

        /**
         * Encodes the specified CreateResponse message. Does not implicitly {@link namespace.CreateResponse.verify|verify} messages.
         * @param message CreateResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.CreateResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CreateResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.CreateResponse;
    }

    /** Properties of a ShowResponse. */
    interface IShowResponse {

        /** ShowResponse namespace */
        namespace?: (types.NamespaceModel|null);
    }

    /** Represents a ShowResponse. */
    class ShowResponse implements IShowResponse {

        /**
         * Constructs a new ShowResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IShowResponse);

        /** ShowResponse namespace. */
        public namespace?: (types.NamespaceModel|null);

        /**
         * Encodes the specified ShowResponse message. Does not implicitly {@link namespace.ShowResponse.verify|verify} messages.
         * @param message ShowResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.ShowResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.ShowResponse;
    }

    /** Properties of an IsExistsResponse. */
    interface IIsExistsResponse {

        /** IsExistsResponse exists */
        exists?: (boolean|null);

        /** IsExistsResponse id */
        id?: (number|null);
    }

    /** Represents an IsExistsResponse. */
    class IsExistsResponse implements IIsExistsResponse {

        /**
         * Constructs a new IsExistsResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IIsExistsResponse);

        /** IsExistsResponse exists. */
        public exists: boolean;

        /** IsExistsResponse id. */
        public id: number;

        /**
         * Encodes the specified IsExistsResponse message. Does not implicitly {@link namespace.IsExistsResponse.verify|verify} messages.
         * @param message IsExistsResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.IsExistsResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an IsExistsResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns IsExistsResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.IsExistsResponse;
    }

    /** Properties of an AllRequest. */
    interface IAllRequest {
    }

    /** Represents an AllRequest. */
    class AllRequest implements IAllRequest {

        /**
         * Constructs a new AllRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IAllRequest);

        /**
         * Encodes the specified AllRequest message. Does not implicitly {@link namespace.AllRequest.verify|verify} messages.
         * @param message AllRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.AllRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.AllRequest;
    }

    /** Properties of a DeleteResponse. */
    interface IDeleteResponse {
    }

    /** Represents a DeleteResponse. */
    class DeleteResponse implements IDeleteResponse {

        /**
         * Constructs a new DeleteResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: namespace.IDeleteResponse);

        /**
         * Encodes the specified DeleteResponse message. Does not implicitly {@link namespace.DeleteResponse.verify|verify} messages.
         * @param message DeleteResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: namespace.DeleteResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): namespace.DeleteResponse;
    }

    /** Represents a Namespace */
    class Namespace extends $protobuf.rpc.Service {

        /**
         * Constructs a new Namespace service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls All.
         * @param request AllRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and AllResponse
         */
        public all(request: namespace.AllRequest, callback: namespace.Namespace.AllCallback): void;

        /**
         * Calls All.
         * @param request AllRequest message or plain object
         * @returns Promise
         */
        public all(request: namespace.AllRequest): Promise<namespace.AllResponse>;

        /**
         * Calls Create.
         * @param request CreateRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and CreateResponse
         */
        public create(request: namespace.CreateRequest, callback: namespace.Namespace.CreateCallback): void;

        /**
         * Calls Create.
         * @param request CreateRequest message or plain object
         * @returns Promise
         */
        public create(request: namespace.CreateRequest): Promise<namespace.CreateResponse>;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ShowResponse
         */
        public show(request: namespace.ShowRequest, callback: namespace.Namespace.ShowCallback): void;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @returns Promise
         */
        public show(request: namespace.ShowRequest): Promise<namespace.ShowResponse>;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DeleteResponse
         */
        public delete(request: namespace.DeleteRequest, callback: namespace.Namespace.DeleteCallback): void;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @returns Promise
         */
        public delete(request: namespace.DeleteRequest): Promise<namespace.DeleteResponse>;

        /**
         * Calls IsExists.
         * @param request IsExistsRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and IsExistsResponse
         */
        public isExists(request: namespace.IsExistsRequest, callback: namespace.Namespace.IsExistsCallback): void;

        /**
         * Calls IsExists.
         * @param request IsExistsRequest message or plain object
         * @returns Promise
         */
        public isExists(request: namespace.IsExistsRequest): Promise<namespace.IsExistsResponse>;
    }

    namespace Namespace {

        /**
         * Callback as used by {@link namespace.Namespace#all}.
         * @param error Error, if any
         * @param [response] AllResponse
         */
        type AllCallback = (error: (Error|null), response?: namespace.AllResponse) => void;

        /**
         * Callback as used by {@link namespace.Namespace#create}.
         * @param error Error, if any
         * @param [response] CreateResponse
         */
        type CreateCallback = (error: (Error|null), response?: namespace.CreateResponse) => void;

        /**
         * Callback as used by {@link namespace.Namespace#show}.
         * @param error Error, if any
         * @param [response] ShowResponse
         */
        type ShowCallback = (error: (Error|null), response?: namespace.ShowResponse) => void;

        /**
         * Callback as used by {@link namespace.Namespace#delete_}.
         * @param error Error, if any
         * @param [response] DeleteResponse
         */
        type DeleteCallback = (error: (Error|null), response?: namespace.DeleteResponse) => void;

        /**
         * Callback as used by {@link namespace.Namespace#isExists}.
         * @param error Error, if any
         * @param [response] IsExistsResponse
         */
        type IsExistsCallback = (error: (Error|null), response?: namespace.IsExistsResponse) => void;
    }
}

/** Namespace picture. */
export namespace picture {

    /** Properties of a BackgroundRequest. */
    interface IBackgroundRequest {

        /** BackgroundRequest random */
        random?: (boolean|null);
    }

    /** Represents a BackgroundRequest. */
    class BackgroundRequest implements IBackgroundRequest {

        /**
         * Constructs a new BackgroundRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: picture.IBackgroundRequest);

        /** BackgroundRequest random. */
        public random: boolean;

        /**
         * Encodes the specified BackgroundRequest message. Does not implicitly {@link picture.BackgroundRequest.verify|verify} messages.
         * @param message BackgroundRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: picture.BackgroundRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BackgroundRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BackgroundRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): picture.BackgroundRequest;
    }

    /** Properties of a BackgroundResponse. */
    interface IBackgroundResponse {

        /** BackgroundResponse url */
        url?: (string|null);

        /** BackgroundResponse copyright */
        copyright?: (string|null);
    }

    /** Represents a BackgroundResponse. */
    class BackgroundResponse implements IBackgroundResponse {

        /**
         * Constructs a new BackgroundResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: picture.IBackgroundResponse);

        /** BackgroundResponse url. */
        public url: string;

        /** BackgroundResponse copyright. */
        public copyright: string;

        /**
         * Encodes the specified BackgroundResponse message. Does not implicitly {@link picture.BackgroundResponse.verify|verify} messages.
         * @param message BackgroundResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: picture.BackgroundResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BackgroundResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BackgroundResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): picture.BackgroundResponse;
    }

    /** Represents a Picture */
    class Picture extends $protobuf.rpc.Service {

        /**
         * Constructs a new Picture service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls Background.
         * @param request BackgroundRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and BackgroundResponse
         */
        public background(request: picture.BackgroundRequest, callback: picture.Picture.BackgroundCallback): void;

        /**
         * Calls Background.
         * @param request BackgroundRequest message or plain object
         * @returns Promise
         */
        public background(request: picture.BackgroundRequest): Promise<picture.BackgroundResponse>;
    }

    namespace Picture {

        /**
         * Callback as used by {@link picture.Picture#background}.
         * @param error Error, if any
         * @param [response] BackgroundResponse
         */
        type BackgroundCallback = (error: (Error|null), response?: picture.BackgroundResponse) => void;
    }
}

/** Namespace project. */
export namespace project {

    /** Properties of a DeleteRequest. */
    interface IDeleteRequest {

        /** DeleteRequest project_id */
        project_id?: (number|null);
    }

    /** Represents a DeleteRequest. */
    class DeleteRequest implements IDeleteRequest {

        /**
         * Constructs a new DeleteRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IDeleteRequest);

        /** DeleteRequest project_id. */
        public project_id: number;

        /**
         * Encodes the specified DeleteRequest message. Does not implicitly {@link project.DeleteRequest.verify|verify} messages.
         * @param message DeleteRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.DeleteRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.DeleteRequest;
    }

    /** Properties of a ShowRequest. */
    interface IShowRequest {

        /** ShowRequest project_id */
        project_id?: (number|null);
    }

    /** Represents a ShowRequest. */
    class ShowRequest implements IShowRequest {

        /**
         * Constructs a new ShowRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IShowRequest);

        /** ShowRequest project_id. */
        public project_id: number;

        /**
         * Encodes the specified ShowRequest message. Does not implicitly {@link project.ShowRequest.verify|verify} messages.
         * @param message ShowRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ShowRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ShowRequest;
    }

    /** Properties of a ShowResponse. */
    interface IShowResponse {

        /** ShowResponse project */
        project?: (types.ProjectModel|null);

        /** ShowResponse urls */
        urls?: (types.ServiceEndpoint[]|null);

        /** ShowResponse cpu */
        cpu?: (string|null);

        /** ShowResponse memory */
        memory?: (string|null);

        /** ShowResponse elements */
        elements?: (mars.Element[]|null);
    }

    /** Represents a ShowResponse. */
    class ShowResponse implements IShowResponse {

        /**
         * Constructs a new ShowResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IShowResponse);

        /** ShowResponse project. */
        public project?: (types.ProjectModel|null);

        /** ShowResponse urls. */
        public urls: types.ServiceEndpoint[];

        /** ShowResponse cpu. */
        public cpu: string;

        /** ShowResponse memory. */
        public memory: string;

        /** ShowResponse elements. */
        public elements: mars.Element[];

        /**
         * Encodes the specified ShowResponse message. Does not implicitly {@link project.ShowResponse.verify|verify} messages.
         * @param message ShowResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ShowResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ShowResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ShowResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ShowResponse;
    }

    /** Properties of an AllContainersRequest. */
    interface IAllContainersRequest {

        /** AllContainersRequest project_id */
        project_id?: (number|null);
    }

    /** Represents an AllContainersRequest. */
    class AllContainersRequest implements IAllContainersRequest {

        /**
         * Constructs a new AllContainersRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IAllContainersRequest);

        /** AllContainersRequest project_id. */
        public project_id: number;

        /**
         * Encodes the specified AllContainersRequest message. Does not implicitly {@link project.AllContainersRequest.verify|verify} messages.
         * @param message AllContainersRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.AllContainersRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllContainersRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllContainersRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.AllContainersRequest;
    }

    /** Properties of an AllContainersResponse. */
    interface IAllContainersResponse {

        /** AllContainersResponse items */
        items?: (types.Container[]|null);
    }

    /** Represents an AllContainersResponse. */
    class AllContainersResponse implements IAllContainersResponse {

        /**
         * Constructs a new AllContainersResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IAllContainersResponse);

        /** AllContainersResponse items. */
        public items: types.Container[];

        /**
         * Encodes the specified AllContainersResponse message. Does not implicitly {@link project.AllContainersResponse.verify|verify} messages.
         * @param message AllContainersResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.AllContainersResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AllContainersResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AllContainersResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.AllContainersResponse;
    }

    /** Properties of an ApplyResponse. */
    interface IApplyResponse {

        /** ApplyResponse metadata */
        metadata?: (websocket.Metadata|null);

        /** ApplyResponse project */
        project?: (types.ProjectModel|null);
    }

    /** Represents an ApplyResponse. */
    class ApplyResponse implements IApplyResponse {

        /**
         * Constructs a new ApplyResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IApplyResponse);

        /** ApplyResponse metadata. */
        public metadata?: (websocket.Metadata|null);

        /** ApplyResponse project. */
        public project?: (types.ProjectModel|null);

        /**
         * Encodes the specified ApplyResponse message. Does not implicitly {@link project.ApplyResponse.verify|verify} messages.
         * @param message ApplyResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ApplyResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ApplyResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ApplyResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ApplyResponse;
    }

    /** Properties of a DryRunApplyResponse. */
    interface IDryRunApplyResponse {

        /** DryRunApplyResponse results */
        results?: (string[]|null);
    }

    /** Represents a DryRunApplyResponse. */
    class DryRunApplyResponse implements IDryRunApplyResponse {

        /**
         * Constructs a new DryRunApplyResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IDryRunApplyResponse);

        /** DryRunApplyResponse results. */
        public results: string[];

        /**
         * Encodes the specified DryRunApplyResponse message. Does not implicitly {@link project.DryRunApplyResponse.verify|verify} messages.
         * @param message DryRunApplyResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.DryRunApplyResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DryRunApplyResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DryRunApplyResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.DryRunApplyResponse;
    }

    /** Properties of an ApplyRequest. */
    interface IApplyRequest {

        /** ApplyRequest namespace_id */
        namespace_id?: (number|null);

        /** ApplyRequest name */
        name?: (string|null);

        /** ApplyRequest git_project_id */
        git_project_id?: (number|null);

        /** ApplyRequest git_branch */
        git_branch?: (string|null);

        /** ApplyRequest git_commit */
        git_commit?: (string|null);

        /** ApplyRequest config */
        config?: (string|null);

        /** ApplyRequest atomic */
        atomic?: (boolean|null);

        /** ApplyRequest websocket_sync */
        websocket_sync?: (boolean|null);

        /** ApplyRequest send_percent */
        send_percent?: (boolean|null);

        /** ApplyRequest extra_values */
        extra_values?: (types.ExtraValue[]|null);

        /** ApplyRequest install_timeout_seconds */
        install_timeout_seconds?: (number|null);
    }

    /** Represents an ApplyRequest. */
    class ApplyRequest implements IApplyRequest {

        /**
         * Constructs a new ApplyRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IApplyRequest);

        /** ApplyRequest namespace_id. */
        public namespace_id: number;

        /** ApplyRequest name. */
        public name: string;

        /** ApplyRequest git_project_id. */
        public git_project_id: number;

        /** ApplyRequest git_branch. */
        public git_branch: string;

        /** ApplyRequest git_commit. */
        public git_commit: string;

        /** ApplyRequest config. */
        public config: string;

        /** ApplyRequest atomic. */
        public atomic: boolean;

        /** ApplyRequest websocket_sync. */
        public websocket_sync: boolean;

        /** ApplyRequest send_percent. */
        public send_percent: boolean;

        /** ApplyRequest extra_values. */
        public extra_values: types.ExtraValue[];

        /** ApplyRequest install_timeout_seconds. */
        public install_timeout_seconds: number;

        /**
         * Encodes the specified ApplyRequest message. Does not implicitly {@link project.ApplyRequest.verify|verify} messages.
         * @param message ApplyRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ApplyRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ApplyRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ApplyRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ApplyRequest;
    }

    /** Properties of a DeleteResponse. */
    interface IDeleteResponse {
    }

    /** Represents a DeleteResponse. */
    class DeleteResponse implements IDeleteResponse {

        /**
         * Constructs a new DeleteResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IDeleteResponse);

        /**
         * Encodes the specified DeleteResponse message. Does not implicitly {@link project.DeleteResponse.verify|verify} messages.
         * @param message DeleteResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.DeleteResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a DeleteResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns DeleteResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.DeleteResponse;
    }

    /** Properties of a ListRequest. */
    interface IListRequest {

        /** ListRequest page */
        page?: (number|null);

        /** ListRequest page_size */
        page_size?: (number|null);
    }

    /** Represents a ListRequest. */
    class ListRequest implements IListRequest {

        /**
         * Constructs a new ListRequest.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IListRequest);

        /** ListRequest page. */
        public page: number;

        /** ListRequest page_size. */
        public page_size: number;

        /**
         * Encodes the specified ListRequest message. Does not implicitly {@link project.ListRequest.verify|verify} messages.
         * @param message ListRequest message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ListRequest, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListRequest message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ListRequest;
    }

    /** Properties of a ListResponse. */
    interface IListResponse {

        /** ListResponse page */
        page?: (number|null);

        /** ListResponse page_size */
        page_size?: (number|null);

        /** ListResponse count */
        count?: (number|null);

        /** ListResponse items */
        items?: (types.ProjectModel[]|null);
    }

    /** Represents a ListResponse. */
    class ListResponse implements IListResponse {

        /**
         * Constructs a new ListResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: project.IListResponse);

        /** ListResponse page. */
        public page: number;

        /** ListResponse page_size. */
        public page_size: number;

        /** ListResponse count. */
        public count: number;

        /** ListResponse items. */
        public items: types.ProjectModel[];

        /**
         * Encodes the specified ListResponse message. Does not implicitly {@link project.ListResponse.verify|verify} messages.
         * @param message ListResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: project.ListResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ListResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ListResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): project.ListResponse;
    }

    /** Represents a Project */
    class Project extends $protobuf.rpc.Service {

        /**
         * Constructs a new Project service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ListResponse
         */
        public list(request: project.ListRequest, callback: project.Project.ListCallback): void;

        /**
         * Calls List.
         * @param request ListRequest message or plain object
         * @returns Promise
         */
        public list(request: project.ListRequest): Promise<project.ListResponse>;

        /**
         * Calls Apply.
         * @param request ApplyRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ApplyResponse
         */
        public apply(request: project.ApplyRequest, callback: project.Project.ApplyCallback): void;

        /**
         * Calls Apply.
         * @param request ApplyRequest message or plain object
         * @returns Promise
         */
        public apply(request: project.ApplyRequest): Promise<project.ApplyResponse>;

        /**
         * Calls ApplyDryRun.
         * @param request ApplyRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DryRunApplyResponse
         */
        public applyDryRun(request: project.ApplyRequest, callback: project.Project.ApplyDryRunCallback): void;

        /**
         * Calls ApplyDryRun.
         * @param request ApplyRequest message or plain object
         * @returns Promise
         */
        public applyDryRun(request: project.ApplyRequest): Promise<project.DryRunApplyResponse>;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and ShowResponse
         */
        public show(request: project.ShowRequest, callback: project.Project.ShowCallback): void;

        /**
         * Calls Show.
         * @param request ShowRequest message or plain object
         * @returns Promise
         */
        public show(request: project.ShowRequest): Promise<project.ShowResponse>;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and DeleteResponse
         */
        public delete(request: project.DeleteRequest, callback: project.Project.DeleteCallback): void;

        /**
         * Calls Delete.
         * @param request DeleteRequest message or plain object
         * @returns Promise
         */
        public delete(request: project.DeleteRequest): Promise<project.DeleteResponse>;

        /**
         * Calls AllContainers.
         * @param request AllContainersRequest message or plain object
         * @param callback Node-style callback called with the error, if any, and AllContainersResponse
         */
        public allContainers(request: project.AllContainersRequest, callback: project.Project.AllContainersCallback): void;

        /**
         * Calls AllContainers.
         * @param request AllContainersRequest message or plain object
         * @returns Promise
         */
        public allContainers(request: project.AllContainersRequest): Promise<project.AllContainersResponse>;
    }

    namespace Project {

        /**
         * Callback as used by {@link project.Project#list}.
         * @param error Error, if any
         * @param [response] ListResponse
         */
        type ListCallback = (error: (Error|null), response?: project.ListResponse) => void;

        /**
         * Callback as used by {@link project.Project#apply}.
         * @param error Error, if any
         * @param [response] ApplyResponse
         */
        type ApplyCallback = (error: (Error|null), response?: project.ApplyResponse) => void;

        /**
         * Callback as used by {@link project.Project#applyDryRun}.
         * @param error Error, if any
         * @param [response] DryRunApplyResponse
         */
        type ApplyDryRunCallback = (error: (Error|null), response?: project.DryRunApplyResponse) => void;

        /**
         * Callback as used by {@link project.Project#show}.
         * @param error Error, if any
         * @param [response] ShowResponse
         */
        type ShowCallback = (error: (Error|null), response?: project.ShowResponse) => void;

        /**
         * Callback as used by {@link project.Project#delete_}.
         * @param error Error, if any
         * @param [response] DeleteResponse
         */
        type DeleteCallback = (error: (Error|null), response?: project.DeleteResponse) => void;

        /**
         * Callback as used by {@link project.Project#allContainers}.
         * @param error Error, if any
         * @param [response] AllContainersResponse
         */
        type AllContainersCallback = (error: (Error|null), response?: project.AllContainersResponse) => void;
    }
}

/** Namespace types. */
export namespace types {

    /** EventActionType enum. */
    enum EventActionType {
        Unknown = 0,
        Create = 1,
        Update = 2,
        Delete = 3,
        Upload = 4,
        Download = 5,
        DryRun = 6,
        Shell = 7
    }

    /** Properties of a Pod. */
    interface IPod {

        /** Pod namespace */
        namespace?: (string|null);

        /** Pod pod */
        pod?: (string|null);
    }

    /** Represents a Pod. */
    class Pod implements IPod {

        /**
         * Constructs a new Pod.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IPod);

        /** Pod namespace. */
        public namespace: string;

        /** Pod pod. */
        public pod: string;

        /**
         * Encodes the specified Pod message. Does not implicitly {@link types.Pod.verify|verify} messages.
         * @param message Pod message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.Pod, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Pod message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Pod
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.Pod;
    }

    /** Properties of a Container. */
    interface IContainer {

        /** Container namespace */
        namespace?: (string|null);

        /** Container pod */
        pod?: (string|null);

        /** Container container */
        container?: (string|null);
    }

    /** Represents a Container. */
    class Container implements IContainer {

        /**
         * Constructs a new Container.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IContainer);

        /** Container namespace. */
        public namespace: string;

        /** Container pod. */
        public pod: string;

        /** Container container. */
        public container: string;

        /**
         * Encodes the specified Container message. Does not implicitly {@link types.Container.verify|verify} messages.
         * @param message Container message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.Container, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Container message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Container
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.Container;
    }

    /** Properties of an ExtraValue. */
    interface IExtraValue {

        /** ExtraValue path */
        path?: (string|null);

        /** ExtraValue value */
        value?: (string|null);
    }

    /** Represents an ExtraValue. */
    class ExtraValue implements IExtraValue {

        /**
         * Constructs a new ExtraValue.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IExtraValue);

        /** ExtraValue path. */
        public path: string;

        /** ExtraValue value. */
        public value: string;

        /**
         * Encodes the specified ExtraValue message. Does not implicitly {@link types.ExtraValue.verify|verify} messages.
         * @param message ExtraValue message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.ExtraValue, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ExtraValue message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ExtraValue
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.ExtraValue;
    }

    /** Properties of a ServiceEndpoint. */
    interface IServiceEndpoint {

        /** ServiceEndpoint name */
        name?: (string|null);

        /** ServiceEndpoint url */
        url?: (string|null);

        /** ServiceEndpoint port_name */
        port_name?: (string|null);
    }

    /** Represents a ServiceEndpoint. */
    class ServiceEndpoint implements IServiceEndpoint {

        /**
         * Constructs a new ServiceEndpoint.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IServiceEndpoint);

        /** ServiceEndpoint name. */
        public name: string;

        /** ServiceEndpoint url. */
        public url: string;

        /** ServiceEndpoint port_name. */
        public port_name: string;

        /**
         * Encodes the specified ServiceEndpoint message. Does not implicitly {@link types.ServiceEndpoint.verify|verify} messages.
         * @param message ServiceEndpoint message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.ServiceEndpoint, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ServiceEndpoint message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ServiceEndpoint
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.ServiceEndpoint;
    }

    /** Properties of a ChangelogModel. */
    interface IChangelogModel {

        /** ChangelogModel id */
        id?: (number|null);

        /** ChangelogModel version */
        version?: (number|null);

        /** ChangelogModel username */
        username?: (string|null);

        /** ChangelogModel manifest */
        manifest?: (string|null);

        /** ChangelogModel config */
        config?: (string|null);

        /** ChangelogModel config_changed */
        config_changed?: (boolean|null);

        /** ChangelogModel project_id */
        project_id?: (number|null);

        /** ChangelogModel git_project_id */
        git_project_id?: (number|null);

        /** ChangelogModel project */
        project?: (types.ProjectModel|null);

        /** ChangelogModel git_project */
        git_project?: (types.GitProjectModel|null);

        /** ChangelogModel date */
        date?: (string|null);

        /** ChangelogModel created_at */
        created_at?: (string|null);

        /** ChangelogModel updated_at */
        updated_at?: (string|null);

        /** ChangelogModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents a ChangelogModel. */
    class ChangelogModel implements IChangelogModel {

        /**
         * Constructs a new ChangelogModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IChangelogModel);

        /** ChangelogModel id. */
        public id: number;

        /** ChangelogModel version. */
        public version: number;

        /** ChangelogModel username. */
        public username: string;

        /** ChangelogModel manifest. */
        public manifest: string;

        /** ChangelogModel config. */
        public config: string;

        /** ChangelogModel config_changed. */
        public config_changed: boolean;

        /** ChangelogModel project_id. */
        public project_id: number;

        /** ChangelogModel git_project_id. */
        public git_project_id: number;

        /** ChangelogModel project. */
        public project?: (types.ProjectModel|null);

        /** ChangelogModel git_project. */
        public git_project?: (types.GitProjectModel|null);

        /** ChangelogModel date. */
        public date: string;

        /** ChangelogModel created_at. */
        public created_at: string;

        /** ChangelogModel updated_at. */
        public updated_at: string;

        /** ChangelogModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified ChangelogModel message. Does not implicitly {@link types.ChangelogModel.verify|verify} messages.
         * @param message ChangelogModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.ChangelogModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ChangelogModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ChangelogModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.ChangelogModel;
    }

    /** Properties of an EventModel. */
    interface IEventModel {

        /** EventModel id */
        id?: (number|null);

        /** EventModel action */
        action?: (types.EventActionType|null);

        /** EventModel username */
        username?: (string|null);

        /** EventModel message */
        message?: (string|null);

        /** EventModel old */
        old?: (string|null);

        /** EventModel new */
        "new"?: (string|null);

        /** EventModel duration */
        duration?: (string|null);

        /** EventModel file_id */
        file_id?: (number|null);

        /** EventModel file */
        file?: (types.FileModel|null);

        /** EventModel event_at */
        event_at?: (string|null);

        /** EventModel created_at */
        created_at?: (string|null);

        /** EventModel updated_at */
        updated_at?: (string|null);

        /** EventModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents an EventModel. */
    class EventModel implements IEventModel {

        /**
         * Constructs a new EventModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IEventModel);

        /** EventModel id. */
        public id: number;

        /** EventModel action. */
        public action: types.EventActionType;

        /** EventModel username. */
        public username: string;

        /** EventModel message. */
        public message: string;

        /** EventModel old. */
        public old: string;

        /** EventModel new. */
        public new: string;

        /** EventModel duration. */
        public duration: string;

        /** EventModel file_id. */
        public file_id: number;

        /** EventModel file. */
        public file?: (types.FileModel|null);

        /** EventModel event_at. */
        public event_at: string;

        /** EventModel created_at. */
        public created_at: string;

        /** EventModel updated_at. */
        public updated_at: string;

        /** EventModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified EventModel message. Does not implicitly {@link types.EventModel.verify|verify} messages.
         * @param message EventModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.EventModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an EventModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns EventModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.EventModel;
    }

    /** Properties of a FileModel. */
    interface IFileModel {

        /** FileModel id */
        id?: (number|null);

        /** FileModel path */
        path?: (string|null);

        /** FileModel size */
        size?: (number|null);

        /** FileModel username */
        username?: (string|null);

        /** FileModel namespace */
        namespace?: (string|null);

        /** FileModel pod */
        pod?: (string|null);

        /** FileModel container */
        container?: (string|null);

        /** FileModel container_Path */
        container_Path?: (string|null);

        /** FileModel humanize_size */
        humanize_size?: (string|null);

        /** FileModel created_at */
        created_at?: (string|null);

        /** FileModel updated_at */
        updated_at?: (string|null);

        /** FileModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents a FileModel. */
    class FileModel implements IFileModel {

        /**
         * Constructs a new FileModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IFileModel);

        /** FileModel id. */
        public id: number;

        /** FileModel path. */
        public path: string;

        /** FileModel size. */
        public size: number;

        /** FileModel username. */
        public username: string;

        /** FileModel namespace. */
        public namespace: string;

        /** FileModel pod. */
        public pod: string;

        /** FileModel container. */
        public container: string;

        /** FileModel container_Path. */
        public container_Path: string;

        /** FileModel humanize_size. */
        public humanize_size: string;

        /** FileModel created_at. */
        public created_at: string;

        /** FileModel updated_at. */
        public updated_at: string;

        /** FileModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified FileModel message. Does not implicitly {@link types.FileModel.verify|verify} messages.
         * @param message FileModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.FileModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a FileModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns FileModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.FileModel;
    }

    /** Properties of a GitProjectModel. */
    interface IGitProjectModel {

        /** GitProjectModel id */
        id?: (number|null);

        /** GitProjectModel default_branch */
        default_branch?: (string|null);

        /** GitProjectModel name */
        name?: (string|null);

        /** GitProjectModel git_project_id */
        git_project_id?: (number|null);

        /** GitProjectModel enabled */
        enabled?: (boolean|null);

        /** GitProjectModel global_enabled */
        global_enabled?: (boolean|null);

        /** GitProjectModel global_config */
        global_config?: (string|null);

        /** GitProjectModel created_at */
        created_at?: (string|null);

        /** GitProjectModel updated_at */
        updated_at?: (string|null);

        /** GitProjectModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents a GitProjectModel. */
    class GitProjectModel implements IGitProjectModel {

        /**
         * Constructs a new GitProjectModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IGitProjectModel);

        /** GitProjectModel id. */
        public id: number;

        /** GitProjectModel default_branch. */
        public default_branch: string;

        /** GitProjectModel name. */
        public name: string;

        /** GitProjectModel git_project_id. */
        public git_project_id: number;

        /** GitProjectModel enabled. */
        public enabled: boolean;

        /** GitProjectModel global_enabled. */
        public global_enabled: boolean;

        /** GitProjectModel global_config. */
        public global_config: string;

        /** GitProjectModel created_at. */
        public created_at: string;

        /** GitProjectModel updated_at. */
        public updated_at: string;

        /** GitProjectModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified GitProjectModel message. Does not implicitly {@link types.GitProjectModel.verify|verify} messages.
         * @param message GitProjectModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.GitProjectModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GitProjectModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GitProjectModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.GitProjectModel;
    }

    /** Properties of an ImagePullSecret. */
    interface IImagePullSecret {

        /** ImagePullSecret name */
        name?: (string|null);
    }

    /** Represents an ImagePullSecret. */
    class ImagePullSecret implements IImagePullSecret {

        /**
         * Constructs a new ImagePullSecret.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IImagePullSecret);

        /** ImagePullSecret name. */
        public name: string;

        /**
         * Encodes the specified ImagePullSecret message. Does not implicitly {@link types.ImagePullSecret.verify|verify} messages.
         * @param message ImagePullSecret message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.ImagePullSecret, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an ImagePullSecret message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ImagePullSecret
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.ImagePullSecret;
    }

    /** Properties of a NamespaceModel. */
    interface INamespaceModel {

        /** NamespaceModel id */
        id?: (number|null);

        /** NamespaceModel name */
        name?: (string|null);

        /** NamespaceModel ImagePullSecrets */
        ImagePullSecrets?: (types.ImagePullSecret[]|null);

        /** NamespaceModel projects */
        projects?: (types.ProjectModel[]|null);

        /** NamespaceModel created_at */
        created_at?: (string|null);

        /** NamespaceModel updated_at */
        updated_at?: (string|null);

        /** NamespaceModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents a NamespaceModel. */
    class NamespaceModel implements INamespaceModel {

        /**
         * Constructs a new NamespaceModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.INamespaceModel);

        /** NamespaceModel id. */
        public id: number;

        /** NamespaceModel name. */
        public name: string;

        /** NamespaceModel ImagePullSecrets. */
        public ImagePullSecrets: types.ImagePullSecret[];

        /** NamespaceModel projects. */
        public projects: types.ProjectModel[];

        /** NamespaceModel created_at. */
        public created_at: string;

        /** NamespaceModel updated_at. */
        public updated_at: string;

        /** NamespaceModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified NamespaceModel message. Does not implicitly {@link types.NamespaceModel.verify|verify} messages.
         * @param message NamespaceModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.NamespaceModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a NamespaceModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns NamespaceModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.NamespaceModel;
    }

    /** Deploy enum. */
    enum Deploy {
        StatusUnknown = 0,
        StatusDeploying = 1,
        StatusDeployed = 2,
        StatusFailed = 3
    }

    /** Properties of a ProjectModel. */
    interface IProjectModel {

        /** ProjectModel id */
        id?: (number|null);

        /** ProjectModel name */
        name?: (string|null);

        /** ProjectModel git_project_id */
        git_project_id?: (number|null);

        /** ProjectModel git_branch */
        git_branch?: (string|null);

        /** ProjectModel git_commit */
        git_commit?: (string|null);

        /** ProjectModel config */
        config?: (string|null);

        /** ProjectModel override_values */
        override_values?: (string|null);

        /** ProjectModel docker_image */
        docker_image?: (string|null);

        /** ProjectModel pod_selectors */
        pod_selectors?: (string|null);

        /** ProjectModel namespace_id */
        namespace_id?: (number|null);

        /** ProjectModel atomic */
        atomic?: (boolean|null);

        /** ProjectModel env_values */
        env_values?: (string|null);

        /** ProjectModel extra_values */
        extra_values?: (types.ExtraValue[]|null);

        /** ProjectModel final_extra_values */
        final_extra_values?: (string|null);

        /** ProjectModel deploy_status */
        deploy_status?: (types.Deploy|null);

        /** ProjectModel humanize_created_at */
        humanize_created_at?: (string|null);

        /** ProjectModel humanize_updated_at */
        humanize_updated_at?: (string|null);

        /** ProjectModel config_type */
        config_type?: (string|null);

        /** ProjectModel git_commit_web_url */
        git_commit_web_url?: (string|null);

        /** ProjectModel git_commit_title */
        git_commit_title?: (string|null);

        /** ProjectModel git_commit_author */
        git_commit_author?: (string|null);

        /** ProjectModel git_commit_date */
        git_commit_date?: (string|null);

        /** ProjectModel namespace */
        namespace?: (types.NamespaceModel|null);

        /** ProjectModel created_at */
        created_at?: (string|null);

        /** ProjectModel updated_at */
        updated_at?: (string|null);

        /** ProjectModel deleted_at */
        deleted_at?: (string|null);
    }

    /** Represents a ProjectModel. */
    class ProjectModel implements IProjectModel {

        /**
         * Constructs a new ProjectModel.
         * @param [properties] Properties to set
         */
        constructor(properties?: types.IProjectModel);

        /** ProjectModel id. */
        public id: number;

        /** ProjectModel name. */
        public name: string;

        /** ProjectModel git_project_id. */
        public git_project_id: number;

        /** ProjectModel git_branch. */
        public git_branch: string;

        /** ProjectModel git_commit. */
        public git_commit: string;

        /** ProjectModel config. */
        public config: string;

        /** ProjectModel override_values. */
        public override_values: string;

        /** ProjectModel docker_image. */
        public docker_image: string;

        /** ProjectModel pod_selectors. */
        public pod_selectors: string;

        /** ProjectModel namespace_id. */
        public namespace_id: number;

        /** ProjectModel atomic. */
        public atomic: boolean;

        /** ProjectModel env_values. */
        public env_values: string;

        /** ProjectModel extra_values. */
        public extra_values: types.ExtraValue[];

        /** ProjectModel final_extra_values. */
        public final_extra_values: string;

        /** ProjectModel deploy_status. */
        public deploy_status: types.Deploy;

        /** ProjectModel humanize_created_at. */
        public humanize_created_at: string;

        /** ProjectModel humanize_updated_at. */
        public humanize_updated_at: string;

        /** ProjectModel config_type. */
        public config_type: string;

        /** ProjectModel git_commit_web_url. */
        public git_commit_web_url: string;

        /** ProjectModel git_commit_title. */
        public git_commit_title: string;

        /** ProjectModel git_commit_author. */
        public git_commit_author: string;

        /** ProjectModel git_commit_date. */
        public git_commit_date: string;

        /** ProjectModel namespace. */
        public namespace?: (types.NamespaceModel|null);

        /** ProjectModel created_at. */
        public created_at: string;

        /** ProjectModel updated_at. */
        public updated_at: string;

        /** ProjectModel deleted_at. */
        public deleted_at: string;

        /**
         * Encodes the specified ProjectModel message. Does not implicitly {@link types.ProjectModel.verify|verify} messages.
         * @param message ProjectModel message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: types.ProjectModel, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a ProjectModel message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns ProjectModel
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): types.ProjectModel;
    }
}

/** Namespace version. */
export namespace version {

    /** Properties of a Request. */
    interface IRequest {
    }

    /** Represents a Request. */
    class Request implements IRequest {

        /**
         * Constructs a new Request.
         * @param [properties] Properties to set
         */
        constructor(properties?: version.IRequest);

        /**
         * Encodes the specified Request message. Does not implicitly {@link version.Request.verify|verify} messages.
         * @param message Request message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: version.Request, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Request message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Request
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): version.Request;
    }

    /** Properties of a Response. */
    interface IResponse {

        /** Response version */
        version?: (string|null);

        /** Response build_date */
        build_date?: (string|null);

        /** Response git_branch */
        git_branch?: (string|null);

        /** Response git_commit */
        git_commit?: (string|null);

        /** Response git_tag */
        git_tag?: (string|null);

        /** Response go_version */
        go_version?: (string|null);

        /** Response compiler */
        compiler?: (string|null);

        /** Response platform */
        platform?: (string|null);

        /** Response kubectl_version */
        kubectl_version?: (string|null);

        /** Response helm_version */
        helm_version?: (string|null);

        /** Response git_repo */
        git_repo?: (string|null);
    }

    /** Represents a Response. */
    class Response implements IResponse {

        /**
         * Constructs a new Response.
         * @param [properties] Properties to set
         */
        constructor(properties?: version.IResponse);

        /** Response version. */
        public version: string;

        /** Response build_date. */
        public build_date: string;

        /** Response git_branch. */
        public git_branch: string;

        /** Response git_commit. */
        public git_commit: string;

        /** Response git_tag. */
        public git_tag: string;

        /** Response go_version. */
        public go_version: string;

        /** Response compiler. */
        public compiler: string;

        /** Response platform. */
        public platform: string;

        /** Response kubectl_version. */
        public kubectl_version: string;

        /** Response helm_version. */
        public helm_version: string;

        /** Response git_repo. */
        public git_repo: string;

        /**
         * Encodes the specified Response message. Does not implicitly {@link version.Response.verify|verify} messages.
         * @param message Response message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: version.Response, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Response message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Response
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): version.Response;
    }

    /** Represents a Version */
    class Version extends $protobuf.rpc.Service {

        /**
         * Constructs a new Version service.
         * @param rpcImpl RPC implementation
         * @param [requestDelimited=false] Whether requests are length-delimited
         * @param [responseDelimited=false] Whether responses are length-delimited
         */
        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

        /**
         * Calls Version.
         * @param request Request message or plain object
         * @param callback Node-style callback called with the error, if any, and Response
         */
        public version(request: version.Request, callback: version.Version.VersionCallback): void;

        /**
         * Calls Version.
         * @param request Request message or plain object
         * @returns Promise
         */
        public version(request: version.Request): Promise<version.Response>;
    }

    namespace Version {

        /**
         * Callback as used by {@link version.Version#version}.
         * @param error Error, if any
         * @param [response] Response
         */
        type VersionCallback = (error: (Error|null), response?: version.Response) => void;
    }
}

/** Namespace websocket. */
export namespace websocket {

    /** Type enum. */
    enum Type {
        TypeUnknown = 0,
        SetUid = 1,
        ReloadProjects = 2,
        CancelProject = 3,
        CreateProject = 4,
        UpdateProject = 5,
        ProcessPercent = 6,
        ClusterInfoSync = 7,
        InternalError = 8,
        ApplyProject = 9,
        HandleExecShell = 50,
        HandleExecShellMsg = 51,
        HandleCloseShell = 52,
        HandleAuthorize = 53
    }

    /** ResultType enum. */
    enum ResultType {
        ResultUnknown = 0,
        Error = 1,
        Success = 2,
        Deployed = 3,
        DeployedFailed = 4,
        DeployedCanceled = 5
    }

    /** To enum. */
    enum To {
        ToSelf = 0,
        ToAll = 1,
        ToOthers = 2
    }

    /** Properties of a WsRequestMetadata. */
    interface IWsRequestMetadata {

        /** WsRequestMetadata type */
        type?: (websocket.Type|null);
    }

    /** Represents a WsRequestMetadata. */
    class WsRequestMetadata implements IWsRequestMetadata {

        /**
         * Constructs a new WsRequestMetadata.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IWsRequestMetadata);

        /** WsRequestMetadata type. */
        public type: websocket.Type;

        /**
         * Encodes the specified WsRequestMetadata message. Does not implicitly {@link websocket.WsRequestMetadata.verify|verify} messages.
         * @param message WsRequestMetadata message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.WsRequestMetadata, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a WsRequestMetadata message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns WsRequestMetadata
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.WsRequestMetadata;
    }

    /** Properties of an AuthorizeTokenInput. */
    interface IAuthorizeTokenInput {

        /** AuthorizeTokenInput type */
        type?: (websocket.Type|null);

        /** AuthorizeTokenInput token */
        token?: (string|null);
    }

    /** Represents an AuthorizeTokenInput. */
    class AuthorizeTokenInput implements IAuthorizeTokenInput {

        /**
         * Constructs a new AuthorizeTokenInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IAuthorizeTokenInput);

        /** AuthorizeTokenInput type. */
        public type: websocket.Type;

        /** AuthorizeTokenInput token. */
        public token: string;

        /**
         * Encodes the specified AuthorizeTokenInput message. Does not implicitly {@link websocket.AuthorizeTokenInput.verify|verify} messages.
         * @param message AuthorizeTokenInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.AuthorizeTokenInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an AuthorizeTokenInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns AuthorizeTokenInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.AuthorizeTokenInput;
    }

    /** Properties of a TerminalMessage. */
    interface ITerminalMessage {

        /** TerminalMessage op */
        op?: (string|null);

        /** TerminalMessage data */
        data?: (string|null);

        /** TerminalMessage session_id */
        session_id?: (string|null);

        /** TerminalMessage rows */
        rows?: (number|null);

        /** TerminalMessage cols */
        cols?: (number|null);
    }

    /** Represents a TerminalMessage. */
    class TerminalMessage implements ITerminalMessage {

        /**
         * Constructs a new TerminalMessage.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.ITerminalMessage);

        /** TerminalMessage op. */
        public op: string;

        /** TerminalMessage data. */
        public data: string;

        /** TerminalMessage session_id. */
        public session_id: string;

        /** TerminalMessage rows. */
        public rows: number;

        /** TerminalMessage cols. */
        public cols: number;

        /**
         * Encodes the specified TerminalMessage message. Does not implicitly {@link websocket.TerminalMessage.verify|verify} messages.
         * @param message TerminalMessage message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.TerminalMessage, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TerminalMessage message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TerminalMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.TerminalMessage;
    }

    /** Properties of a TerminalMessageInput. */
    interface ITerminalMessageInput {

        /** TerminalMessageInput type */
        type?: (websocket.Type|null);

        /** TerminalMessageInput message */
        message?: (websocket.TerminalMessage|null);
    }

    /** Represents a TerminalMessageInput. */
    class TerminalMessageInput implements ITerminalMessageInput {

        /**
         * Constructs a new TerminalMessageInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.ITerminalMessageInput);

        /** TerminalMessageInput type. */
        public type: websocket.Type;

        /** TerminalMessageInput message. */
        public message?: (websocket.TerminalMessage|null);

        /**
         * Encodes the specified TerminalMessageInput message. Does not implicitly {@link websocket.TerminalMessageInput.verify|verify} messages.
         * @param message TerminalMessageInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.TerminalMessageInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a TerminalMessageInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns TerminalMessageInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.TerminalMessageInput;
    }

    /** Properties of a WsHandleExecShellInput. */
    interface IWsHandleExecShellInput {

        /** WsHandleExecShellInput type */
        type?: (websocket.Type|null);

        /** WsHandleExecShellInput container */
        container?: (types.Container|null);
    }

    /** Represents a WsHandleExecShellInput. */
    class WsHandleExecShellInput implements IWsHandleExecShellInput {

        /**
         * Constructs a new WsHandleExecShellInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IWsHandleExecShellInput);

        /** WsHandleExecShellInput type. */
        public type: websocket.Type;

        /** WsHandleExecShellInput container. */
        public container?: (types.Container|null);

        /**
         * Encodes the specified WsHandleExecShellInput message. Does not implicitly {@link websocket.WsHandleExecShellInput.verify|verify} messages.
         * @param message WsHandleExecShellInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.WsHandleExecShellInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a WsHandleExecShellInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns WsHandleExecShellInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.WsHandleExecShellInput;
    }

    /** Properties of a CancelInput. */
    interface ICancelInput {

        /** CancelInput type */
        type?: (websocket.Type|null);

        /** CancelInput namespace_id */
        namespace_id?: (number|null);

        /** CancelInput name */
        name?: (string|null);
    }

    /** Represents a CancelInput. */
    class CancelInput implements ICancelInput {

        /**
         * Constructs a new CancelInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.ICancelInput);

        /** CancelInput type. */
        public type: websocket.Type;

        /** CancelInput namespace_id. */
        public namespace_id: number;

        /** CancelInput name. */
        public name: string;

        /**
         * Encodes the specified CancelInput message. Does not implicitly {@link websocket.CancelInput.verify|verify} messages.
         * @param message CancelInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.CancelInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CancelInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CancelInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.CancelInput;
    }

    /** Properties of a CreateProjectInput. */
    interface ICreateProjectInput {

        /** CreateProjectInput type */
        type?: (websocket.Type|null);

        /** CreateProjectInput namespace_id */
        namespace_id?: (number|null);

        /** CreateProjectInput name */
        name?: (string|null);

        /** CreateProjectInput git_project_id */
        git_project_id?: (number|null);

        /** CreateProjectInput git_branch */
        git_branch?: (string|null);

        /** CreateProjectInput git_commit */
        git_commit?: (string|null);

        /** CreateProjectInput config */
        config?: (string|null);

        /** CreateProjectInput atomic */
        atomic?: (boolean|null);

        /** CreateProjectInput extra_values */
        extra_values?: (types.ExtraValue[]|null);
    }

    /** Represents a CreateProjectInput. */
    class CreateProjectInput implements ICreateProjectInput {

        /**
         * Constructs a new CreateProjectInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.ICreateProjectInput);

        /** CreateProjectInput type. */
        public type: websocket.Type;

        /** CreateProjectInput namespace_id. */
        public namespace_id: number;

        /** CreateProjectInput name. */
        public name: string;

        /** CreateProjectInput git_project_id. */
        public git_project_id: number;

        /** CreateProjectInput git_branch. */
        public git_branch: string;

        /** CreateProjectInput git_commit. */
        public git_commit: string;

        /** CreateProjectInput config. */
        public config: string;

        /** CreateProjectInput atomic. */
        public atomic: boolean;

        /** CreateProjectInput extra_values. */
        public extra_values: types.ExtraValue[];

        /**
         * Encodes the specified CreateProjectInput message. Does not implicitly {@link websocket.CreateProjectInput.verify|verify} messages.
         * @param message CreateProjectInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.CreateProjectInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a CreateProjectInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns CreateProjectInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.CreateProjectInput;
    }

    /** Properties of an UpdateProjectInput. */
    interface IUpdateProjectInput {

        /** UpdateProjectInput type */
        type?: (websocket.Type|null);

        /** UpdateProjectInput project_id */
        project_id?: (number|null);

        /** UpdateProjectInput git_branch */
        git_branch?: (string|null);

        /** UpdateProjectInput git_commit */
        git_commit?: (string|null);

        /** UpdateProjectInput config */
        config?: (string|null);

        /** UpdateProjectInput atomic */
        atomic?: (boolean|null);

        /** UpdateProjectInput extra_values */
        extra_values?: (types.ExtraValue[]|null);
    }

    /** Represents an UpdateProjectInput. */
    class UpdateProjectInput implements IUpdateProjectInput {

        /**
         * Constructs a new UpdateProjectInput.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IUpdateProjectInput);

        /** UpdateProjectInput type. */
        public type: websocket.Type;

        /** UpdateProjectInput project_id. */
        public project_id: number;

        /** UpdateProjectInput git_branch. */
        public git_branch: string;

        /** UpdateProjectInput git_commit. */
        public git_commit: string;

        /** UpdateProjectInput config. */
        public config: string;

        /** UpdateProjectInput atomic. */
        public atomic: boolean;

        /** UpdateProjectInput extra_values. */
        public extra_values: types.ExtraValue[];

        /**
         * Encodes the specified UpdateProjectInput message. Does not implicitly {@link websocket.UpdateProjectInput.verify|verify} messages.
         * @param message UpdateProjectInput message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.UpdateProjectInput, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes an UpdateProjectInput message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns UpdateProjectInput
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.UpdateProjectInput;
    }

    /** Properties of a Metadata. */
    interface IMetadata {

        /** Metadata id */
        id?: (string|null);

        /** Metadata uid */
        uid?: (string|null);

        /** Metadata slug */
        slug?: (string|null);

        /** Metadata type */
        type?: (websocket.Type|null);

        /** Metadata end */
        end?: (boolean|null);

        /** Metadata result */
        result?: (websocket.ResultType|null);

        /** Metadata to */
        to?: (websocket.To|null);

        /** Metadata message */
        message?: (string|null);
    }

    /** Represents a Metadata. */
    class Metadata implements IMetadata {

        /**
         * Constructs a new Metadata.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IMetadata);

        /** Metadata id. */
        public id: string;

        /** Metadata uid. */
        public uid: string;

        /** Metadata slug. */
        public slug: string;

        /** Metadata type. */
        public type: websocket.Type;

        /** Metadata end. */
        public end: boolean;

        /** Metadata result. */
        public result: websocket.ResultType;

        /** Metadata to. */
        public to: websocket.To;

        /** Metadata message. */
        public message: string;

        /**
         * Encodes the specified Metadata message. Does not implicitly {@link websocket.Metadata.verify|verify} messages.
         * @param message Metadata message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.Metadata, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Metadata message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Metadata
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.Metadata;
    }

    /** Properties of a WsMetadataResponse. */
    interface IWsMetadataResponse {

        /** WsMetadataResponse metadata */
        metadata?: (websocket.Metadata|null);
    }

    /** Represents a WsMetadataResponse. */
    class WsMetadataResponse implements IWsMetadataResponse {

        /**
         * Constructs a new WsMetadataResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IWsMetadataResponse);

        /** WsMetadataResponse metadata. */
        public metadata?: (websocket.Metadata|null);

        /**
         * Encodes the specified WsMetadataResponse message. Does not implicitly {@link websocket.WsMetadataResponse.verify|verify} messages.
         * @param message WsMetadataResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.WsMetadataResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a WsMetadataResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns WsMetadataResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.WsMetadataResponse;
    }

    /** Properties of a WsHandleShellResponse. */
    interface IWsHandleShellResponse {

        /** WsHandleShellResponse metadata */
        metadata?: (websocket.Metadata|null);

        /** WsHandleShellResponse terminal_message */
        terminal_message?: (websocket.TerminalMessage|null);

        /** WsHandleShellResponse container */
        container?: (types.Container|null);
    }

    /** Represents a WsHandleShellResponse. */
    class WsHandleShellResponse implements IWsHandleShellResponse {

        /**
         * Constructs a new WsHandleShellResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IWsHandleShellResponse);

        /** WsHandleShellResponse metadata. */
        public metadata?: (websocket.Metadata|null);

        /** WsHandleShellResponse terminal_message. */
        public terminal_message?: (websocket.TerminalMessage|null);

        /** WsHandleShellResponse container. */
        public container?: (types.Container|null);

        /**
         * Encodes the specified WsHandleShellResponse message. Does not implicitly {@link websocket.WsHandleShellResponse.verify|verify} messages.
         * @param message WsHandleShellResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.WsHandleShellResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a WsHandleShellResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns WsHandleShellResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.WsHandleShellResponse;
    }

    /** Properties of a WsHandleClusterResponse. */
    interface IWsHandleClusterResponse {

        /** WsHandleClusterResponse metadata */
        metadata?: (websocket.Metadata|null);

        /** WsHandleClusterResponse info */
        info?: (cluster.InfoResponse|null);
    }

    /** Represents a WsHandleClusterResponse. */
    class WsHandleClusterResponse implements IWsHandleClusterResponse {

        /**
         * Constructs a new WsHandleClusterResponse.
         * @param [properties] Properties to set
         */
        constructor(properties?: websocket.IWsHandleClusterResponse);

        /** WsHandleClusterResponse metadata. */
        public metadata?: (websocket.Metadata|null);

        /** WsHandleClusterResponse info. */
        public info?: (cluster.InfoResponse|null);

        /**
         * Encodes the specified WsHandleClusterResponse message. Does not implicitly {@link websocket.WsHandleClusterResponse.verify|verify} messages.
         * @param message WsHandleClusterResponse message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: websocket.WsHandleClusterResponse, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a WsHandleClusterResponse message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns WsHandleClusterResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): websocket.WsHandleClusterResponse;
    }
}
